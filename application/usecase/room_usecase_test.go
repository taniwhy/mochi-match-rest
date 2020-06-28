package usecase

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
	"github.com/taniwhy/mochi-match-rest/domain/service/mock_service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
	"github.com/taniwhy/mochi-match-rest/util/testutil"

	mock_repository "github.com/taniwhy/mochi-match-rest/domain/repository/mock_repository"
)

func TestGetRoomList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoomRepository := mock_repository.NewMockRoomRepository(ctrl)
	mockRoomRepository.EXPECT().FindByLimitAndOffset(8, 0).Return([]*output.RoomResBody{}, nil)

	mockEntryHistoryRepository := mock_repository.NewMockEntryHistoryRepository(ctrl)
	mockRoomService := mock_service.NewMockIRoomService(ctrl)

	test := NewRoomUsecase(mockRoomRepository, mockEntryHistoryRepository, mockRoomService)

	// 正常処理テスト
	req, _ := http.NewRequest("GET", "?page=1", nil)
	context := &gin.Context{Request: req}
	user, err := test.GetList(context)

	assert.NotNil(t, user)
	assert.NoError(t, err)

	// 異常処理テスト
	// 1. query無し
	req, _ = http.NewRequest("GET", "", nil)
	context = &gin.Context{Request: req}
	user, err = test.GetList(context)

	assert.Empty(t, user)
	assert.Error(t, err)

	// 2. queryの異常値
	req, _ = http.NewRequest("GET", "?page=foo", nil)
	context = &gin.Context{Request: req}
	user, err = test.GetList(context)

	assert.Empty(t, user)
	assert.Error(t, err)

	req, _ = http.NewRequest("GET", "?page=", nil)
	context = &gin.Context{Request: req}
	user, err = test.GetList(context)

	assert.Empty(t, user)
	assert.Error(t, err)

	req, _ = http.NewRequest("GET", "?page=-1", nil)
	context = &gin.Context{Request: req}
	user, err = test.GetList(context)

	assert.Empty(t, user)
	assert.Error(t, err)

	req, _ = http.NewRequest("GET", "?page=0", nil)
	context = &gin.Context{Request: req}
	user, err = test.GetList(context)

	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestCreateRoom(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	testDate := time.Date(2020, time.April, 1, 00, 00, 00, 00, jst)
	testutil.SetFakeUuID("testUUID")
	testutil.SetFakeTime(testDate)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	start, _ := time.Parse(time.RFC3339, "2020-04-01T00:00:00+09:00")
	room, _ := models.NewRoom("existID", "testText", "testID", "testID", 4, start)
	entryHistory, _ := models.NewEntryHistory("existID", "testUUID")

	mockRoomRepository := mock_repository.NewMockRoomRepository(ctrl)
	mockRoomRepository.EXPECT().Insert(room).Return(nil)

	mockEntryHistoryRepository := mock_repository.NewMockEntryHistoryRepository(ctrl)
	mockEntryHistoryRepository.EXPECT().Insert(entryHistory).Return(nil)

	mockRoomService := mock_service.NewMockIRoomService(ctrl)
	mockRoomService.EXPECT().CanInsert("existID").Return(true, nil)

	test := NewRoomUsecase(mockRoomRepository, mockEntryHistoryRepository, mockRoomService)

	existToken := auth.GenerateAccessToken("existID", false)
	invalidToken := existToken + "foo"

	// 正常処理テスト
	bodyReader := strings.NewReader(
		`{
			"room_text": "testText",
			"game_hard_id": "testID",
			"game_list_id": "testID",
			"capacity": 4,
			"start": "2020-04-01T00:00:00+09:00"
		}`)
	req, _ := http.NewRequest("GET", "", bodyReader)
	req.Header.Add("Authorization", existToken)
	context := &gin.Context{Request: req}
	err := test.Create(context)

	assert.NoError(t, err)

	// 異常処理テスト
	// 1. トークン無し
	bodyReader = strings.NewReader(
		`{
			"room_text": "testText",
			"game_hard_id": "testID",
			"game_list_id": "testID",
			"capacity": 4,
			"start": "2020-04-01T00:00:00+09:00"
		}`)
	req, _ = http.NewRequest("GET", "", bodyReader)
	context = &gin.Context{Request: req}
	err = test.Create(context)

	assert.Error(t, err)

	// 2. 異常なトークン
	bodyReader = strings.NewReader(
		`{
			"room_text": "testText",
			"game_hard_id": "testID",
			"game_list_id": "testID",
			"capacity": 4,
			"start": "2020-04-01T00:00:00+09:00"
		}`)
	req, _ = http.NewRequest("GET", "", bodyReader)
	req.Header.Add("Authorization", invalidToken)
	context = &gin.Context{Request: req}
	err = test.Create(context)

	assert.Error(t, err)
}
