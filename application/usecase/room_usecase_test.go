package usecase

import (
	"errors"
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

	mockRoomRepository := mock_repository.NewMockIRoomRepository(ctrl)
	mockRoomRepository.EXPECT().FindByLimitAndOffset(8, 0).Return([]*output.RoomResBody{}, nil)

	mockEntryHistoryRepository := mock_repository.NewMockIEntryHistoryRepository(ctrl)
	mockRoomService := mock_service.NewMockIRoomService(ctrl)
	mockEntryHistoryService := mock_service.NewMockIEntryHistoryService(ctrl)

	test := NewRoomUsecase(mockRoomRepository, mockEntryHistoryRepository, mockRoomService, mockEntryHistoryService)

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

	mockRoomRepository := mock_repository.NewMockIRoomRepository(ctrl)
	mockRoomRepository.EXPECT().Insert(room).Return(nil)

	mockEntryHistoryRepository := mock_repository.NewMockIEntryHistoryRepository(ctrl)
	mockEntryHistoryRepository.EXPECT().Insert(entryHistory).Return(nil)

	mockRoomService := mock_service.NewMockIRoomService(ctrl)
	mockRoomService.EXPECT().CanInsert("existID").Return(true, nil)

	mockEntryHistoryService := mock_service.NewMockIEntryHistoryService(ctrl)

	test := NewRoomUsecase(mockRoomRepository, mockEntryHistoryRepository, mockRoomService, mockEntryHistoryService)

	existToken := auth.GenerateAccessToken("existID", false)
	invalidToken := existToken + "foo"

	// 正常処理テスト
	bodyReader := strings.NewReader(`
		{
			"room_text": "testText",
			"game_hard_id": "testID",
			"game_list_id": "testID",
			"capacity": 4,
			"start": "2020-04-01T00:00:00+09:00"
		}
		`)
	req, _ := http.NewRequest("GET", "", bodyReader)
	req.Header.Add("Authorization", existToken)
	context := &gin.Context{Request: req}
	err := test.Create(context)

	assert.NoError(t, err)

	// 異常処理テスト
	// 1. トークン無し
	bodyReader = strings.NewReader(`
		{
			"room_text": "testText",
			"game_hard_id": "testID",
			"game_list_id": "testID",
			"capacity": 4,
			"start": "2020-04-01T00:00:00+09:00"
		}
		`)
	req, _ = http.NewRequest("GET", "", bodyReader)
	context = &gin.Context{Request: req}
	err = test.Create(context)

	assert.Error(t, err)

	// 2. 異常なトークン
	bodyReader = strings.NewReader(`
		{
			"room_text": "testText",
			"game_hard_id": "testID",
			"game_list_id": "testID",
			"capacity": 4,
			"start": "2020-04-01T00:00:00+09:00"
		}
		`)
	req, _ = http.NewRequest("GET", "", bodyReader)
	req.Header.Add("Authorization", invalidToken)
	context = &gin.Context{Request: req}
	err = test.Create(context)

	assert.Error(t, err)
}

func TestUpdateRoom(t *testing.T) {}

func TestDeleteRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoomRepository := mock_repository.NewMockIRoomRepository(ctrl)
	mockRoomRepository.EXPECT().LockFlg("existID", "existRoomID").Return(nil)

	mockEntryHistoryRepository := mock_repository.NewMockIEntryHistoryRepository(ctrl)
	mockRoomService := mock_service.NewMockIRoomService(ctrl)
	mockRoomService.EXPECT().IsOwner("existID", "existRoomID").Return(true, nil)
	mockRoomService.EXPECT().IsOwner("existID", "notExistRoomID").Return(false, errors.New("not exist room"))
	mockRoomService.EXPECT().IsOwner("norRoomOwnerID", "existRoomID").Return(false, nil)

	mockEntryHistoryService := mock_service.NewMockIEntryHistoryService(ctrl)

	test := NewRoomUsecase(mockRoomRepository, mockEntryHistoryRepository, mockRoomService, mockEntryHistoryService)

	existToken := auth.GenerateAccessToken("existID", false)
	norRoomOwnerToken := auth.GenerateAccessToken("norRoomOwnerID", false)
	invalidToken := existToken + "foo"

	// 正常処理テスト
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	param := gin.Param{Key: "id", Value: "existRoomID"}
	params := gin.Params{param}
	context := &gin.Context{Request: req, Params: params}
	err := test.Delete(context)

	assert.NoError(t, err)

	// 異常処理テスト
	// 1. トークン無し
	req, _ = http.NewRequest("GET", "", nil)
	context = &gin.Context{Request: req}
	err = test.Delete(context)

	assert.Error(t, err)

	// 2. 異常なトークン
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", invalidToken)
	context = &gin.Context{Request: req}
	err = test.Delete(context)

	assert.Error(t, err)

	// 3. 存在しないルームID
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	param = gin.Param{Key: "id", Value: "notExistRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Delete(context)

	assert.Error(t, err)

	// 3. 非ルームオーナー
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", norRoomOwnerToken)
	param = gin.Param{Key: "id", Value: "existRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Delete(context)

	assert.Error(t, err)
}

func TestJoinRoom(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	testDate := time.Date(2020, time.April, 1, 00, 00, 00, 00, jst)
	testutil.SetFakeUuID("testUUID")
	testutil.SetFakeTime(testDate)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoomRepository := mock_repository.NewMockIRoomRepository(ctrl)
	mockRoomRepository.EXPECT().FindByID("notLockRoomID").Return(&models.Room{Capacity: 2}, nil)

	history, _ := models.NewEntryHistory("existID", "notLockRoomID")

	mockEntryHistoryRepository := mock_repository.NewMockIEntryHistoryRepository(ctrl)
	mockEntryHistoryRepository.EXPECT().CountEntryUser("notLockRoomID").Return(1, nil)
	mockEntryHistoryRepository.EXPECT().Insert(history).Return(nil)

	mockRoomService := mock_service.NewMockIRoomService(ctrl)
	mockRoomService.EXPECT().IsLock("notLockRoomID").Return(true, nil)
	mockRoomService.EXPECT().IsLock("notLockRoomID").Return(true, nil)
	mockRoomService.EXPECT().IsLock("notLockRoomID").Return(true, nil)
	mockRoomService.EXPECT().IsLock("lockRoomID").Return(false, nil)
	mockRoomService.EXPECT().IsLock("lockRoomID").Return(false, nil)

	mockEntryHistoryService := mock_service.NewMockIEntryHistoryService(ctrl)
	mockEntryHistoryService.EXPECT().CanJoin("existID").Return(true, nil)

	test := NewRoomUsecase(mockRoomRepository, mockEntryHistoryRepository, mockRoomService, mockEntryHistoryService)

	existToken := auth.GenerateAccessToken("existID", false)
	invalidToken := existToken + "foo"

	// 正常処理テスト
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	param := gin.Param{Key: "id", Value: "notLockRoomID"}
	params := gin.Params{param}
	context := &gin.Context{Request: req, Params: params}
	err := test.Join(context)

	assert.NoError(t, err)

	// 異常処理テスト
	// 1. トークン無し
	req, _ = http.NewRequest("GET", "", nil)
	param = gin.Param{Key: "id", Value: "notLockRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Join(context)

	assert.Error(t, err)

	// 2. 異常なトークン
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", invalidToken)
	param = gin.Param{Key: "id", Value: "notLockRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Join(context)

	assert.Error(t, err)

	// 3. Param無し
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	context = &gin.Context{Request: req}
	err = test.Join(context)

	assert.Error(t, err)

	// 4. ロック済みルーム参加
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	param = gin.Param{Key: "id", Value: "lockRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Join(context)

	assert.Error(t, err)

	// 4. 既入室ユーザーのルーム参加
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	param = gin.Param{Key: "id", Value: "lockRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Join(context)

	assert.Error(t, err)
}

func TestLaveRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoomRepository := mock_repository.NewMockIRoomRepository(ctrl)
	mockEntryHistoryRepository := mock_repository.NewMockIEntryHistoryRepository(ctrl)
	mockEntryHistoryRepository.EXPECT().LeaveFlg("existRoomID", "existID").Return(nil)

	mockRoomService := mock_service.NewMockIRoomService(ctrl)
	mockEntryHistoryService := mock_service.NewMockIEntryHistoryService(ctrl)
	mockEntryHistoryService.EXPECT().CheckJoin("existRoomID", "existID").Return(false, nil)

	test := NewRoomUsecase(mockRoomRepository, mockEntryHistoryRepository, mockRoomService, mockEntryHistoryService)

	existToken := auth.GenerateAccessToken("existID", false)
	invalidToken := existToken + "foo"

	// 正常処理テスト
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	param := gin.Param{Key: "id", Value: "existRoomID"}
	params := gin.Params{param}
	context := &gin.Context{Request: req, Params: params}
	err := test.Leave(context)

	assert.NoError(t, err)

	// 異常処理テスト
	// 1. トークン無し
	req, _ = http.NewRequest("GET", "", nil)
	param = gin.Param{Key: "id", Value: "notLockRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Leave(context)

	assert.Error(t, err)

	// 2. 異常なトークン
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", invalidToken)
	param = gin.Param{Key: "id", Value: "notLockRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Leave(context)

	assert.Error(t, err)

	// 3. Param無し
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	context = &gin.Context{Request: req}
	err = test.Leave(context)

	assert.Error(t, err)
}
