package usecase

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/taniwhy/mochi-match-rest/domain/models"
	mock_repository "github.com/taniwhy/mochi-match-rest/domain/repository/mock_repository"
	"github.com/taniwhy/mochi-match-rest/domain/service/mock_service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
	"github.com/taniwhy/mochi-match-rest/util/testutil"
)

func TestGetBlacklistByRoomID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlacklistRepository := mock_repository.NewMockIRoomBlacklistRepository(ctrl)
	mockBlacklistRepository.EXPECT().FindByRoomID("roomID").Return([]*models.RoomBlacklist{}, nil)
	mockBlacklistRepository.EXPECT().FindByRoomID("notExistRoomID").Return(nil, errors.New("not found"))

	mockRoomService := mock_service.NewMockIRoomService(ctrl)

	test := NewRoomBlacklistUsecase(mockBlacklistRepository, mockRoomService)

	// 正常処理テスト
	req, _ := http.NewRequest("GET", "", nil)
	param := gin.Param{Key: "id", Value: "roomID"}
	params := gin.Params{param}
	context := &gin.Context{Request: req, Params: params}
	user, err := test.GetByRoomID(context)

	assert.NotNil(t, user)
	assert.NoError(t, err)

	// 異常処理テスト
	// 1. 存在しないルームID
	req, _ = http.NewRequest("GET", "", nil)
	param = gin.Param{Key: "id", Value: "notExistRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	user, err = test.GetByRoomID(context)

	assert.Nil(t, user)
	assert.Error(t, err)
}

func TestCreateBlacklist(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	testDate := time.Date(2020, time.April, 1, 00, 00, 00, 00, jst)
	testutil.SetFakeUuID("testUUID")
	testutil.SetFakeTime(testDate)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlacklistRepository := mock_repository.NewMockIRoomBlacklistRepository(ctrl)
	blacklist, _ := models.NewBlacklist("existRoomID", "existID")
	mockBlacklistRepository.EXPECT().Insert(blacklist).Return(nil)

	mockRoomService := mock_service.NewMockIRoomService(ctrl)
	mockRoomService.EXPECT().IsOwner("existID", "existRoomID").Return(true, nil)
	mockRoomService.EXPECT().IsOwner("notOwnerID", "existRoomID").Return(false, nil)

	test := NewRoomBlacklistUsecase(mockBlacklistRepository, mockRoomService)

	existToken := auth.GenerateAccessToken("existID", false)
	notOwnerToken := auth.GenerateAccessToken("notOwnerID", false)
	invalidToken := existToken + "foo"

	// 正常処理テスト
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	param := gin.Param{Key: "id", Value: "existRoomID"}
	params := gin.Params{param}
	context := &gin.Context{Request: req, Params: params}
	err := test.Create(context)

	assert.NoError(t, err)

	// 異常処理テスト
	// 1. 非ルームオーナー
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", notOwnerToken)
	param = gin.Param{Key: "id", Value: "existRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Create(context)

	assert.Error(t, err)

	// 2. トークン無し
	req, _ = http.NewRequest("GET", "", nil)
	param = gin.Param{Key: "id", Value: "existRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Create(context)

	assert.Error(t, err)

	// 3. トークンの異常
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", invalidToken)
	param = gin.Param{Key: "id", Value: "existRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Create(context)

	assert.Error(t, err)
}

func TestDeleteBlacklist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlacklistRepository := mock_repository.NewMockIRoomBlacklistRepository(ctrl)
	mockBlacklistRepository.EXPECT().Delete("existRoomID", "existID").Return(nil)

	mockRoomService := mock_service.NewMockIRoomService(ctrl)
	mockRoomService.EXPECT().IsOwner("existID", "existRoomID").Return(true, nil)
	mockRoomService.EXPECT().IsOwner("notOwnerID", "existRoomID").Return(false, nil)

	test := NewRoomBlacklistUsecase(mockBlacklistRepository, mockRoomService)

	existToken := auth.GenerateAccessToken("existID", false)
	notOwnerToken := auth.GenerateAccessToken("notOwnerID", false)
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
	// 1. 非ルームオーナー
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", notOwnerToken)
	param = gin.Param{Key: "id", Value: "existRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Delete(context)

	assert.Error(t, err)

	// 2. トークン無し
	req, _ = http.NewRequest("GET", "", nil)
	param = gin.Param{Key: "id", Value: "existRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Delete(context)

	assert.Error(t, err)

	// 3. トークンの異常
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", invalidToken)
	param = gin.Param{Key: "id", Value: "existRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Delete(context)

	assert.Error(t, err)
}
