package usecase

import (
	"database/sql"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	mock_repository "github.com/taniwhy/mochi-match-rest/domain/repository/mock_repository"
	mock_service "github.com/taniwhy/mochi-match-rest/domain/service/mock_service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
	"github.com/taniwhy/mochi-match-rest/util/testutil"
)

func TestGetMe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindByID("existID").Return(&models.User{}, nil)
	mockUserRepository.EXPECT().FindByID("notExistID").Return(nil, nil)

	mockUserDetailRepository := mock_repository.NewMockUserDetailRepository(ctrl)
	mockUserDetailRepository.EXPECT().FindByID("existID").Return(&models.UserDetail{}, nil)

	mockUserService := mock_service.NewMockIUserService(ctrl)
	mockUserService.EXPECT().IsDelete("existID").Return(true, nil)
	mockUserService.EXPECT().IsDelete("notExistID").Return(false, nil)

	mockFavoriteGameRepository := mock_repository.NewMockFavoriteGameRepository(ctrl)
	mockFavoriteGameRepository.EXPECT().FindByID("existID").Return([]*models.FavoriteGame{{}, {}}, nil)

	test := NewUserUsecase(mockUserRepository, mockUserDetailRepository, mockUserService, mockFavoriteGameRepository)

	existToken := auth.GenerateAccessToken("existID", false)
	notExistToken := auth.GenerateAccessToken("notExistID", false)
	invalidToken := "foo"

	// 正常処理テスト
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", existToken)
	context := &gin.Context{Request: req}
	user, err := test.GetMe(context)

	assert.NotEmpty(t, user)
	assert.NoError(t, err)

	// 異常処理テスト
	// 1. 登録されていないユーザートークン
	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", notExistToken)
	context = &gin.Context{Request: req}
	user, err = test.GetMe(context)

	assert.Empty(t, user)
	assert.Error(t, err)

	// 2. 無効なトークン
	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", invalidToken)
	context = &gin.Context{Request: req}
	user, err = test.GetMe(context)

	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindByID("existID").Return(&models.User{}, nil)

	mockUserDetailRepository := mock_repository.NewMockUserDetailRepository(ctrl)
	mockUserDetailRepository.EXPECT().FindByID("existID").Return(&models.UserDetail{}, nil)

	mockUserService := mock_service.NewMockIUserService(ctrl)
	mockUserService.EXPECT().IsDelete("existID").Return(true, nil)
	mockUserService.EXPECT().IsDelete("notExistID").Return(false, nil)

	mockFavoriteGameRepository := mock_repository.NewMockFavoriteGameRepository(ctrl)
	mockFavoriteGameRepository.EXPECT().FindByID("existID").Return([]*models.FavoriteGame{{}, {}}, nil)

	test := NewUserUsecase(mockUserRepository, mockUserDetailRepository, mockUserService, mockFavoriteGameRepository)

	// 正常処理テスト
	req, _ := http.NewRequest("GET", "", nil)
	param := gin.Param{Key: "id", Value: "existID"}
	params := gin.Params{param}
	context := &gin.Context{Request: req, Params: params}
	user, err := test.GetByID(context)

	assert.NotEmpty(t, user)
	assert.NoError(t, err)

	// 異常処理テスト
	// 1. 登録されていないID
	req, _ = http.NewRequest("GET", "", nil)
	param = gin.Param{Key: "id", Value: "notExistID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	user, err = test.GetByID(context)

	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestGetUserByProviderID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindByProviderID("google", "existID").Return(&models.User{UserID: "foo"}, nil)
	mockUserRepository.EXPECT().FindByProviderID("google", "notExistID").Return(nil, nil)
	mockUserRepository.EXPECT().FindByProviderID("foo", "bar").Return(nil, errors.ErrUnexpectedQueryProvider{})

	mockUserDetailRepository := mock_repository.NewMockUserDetailRepository(ctrl)
	mockUserService := mock_service.NewMockIUserService(ctrl)
	mockFavoriteGameRepository := mock_repository.NewMockFavoriteGameRepository(ctrl)

	test := NewUserUsecase(mockUserRepository, mockUserDetailRepository, mockUserService, mockFavoriteGameRepository)

	// 正常処理テスト
	user, err := test.GetByProviderID("google", "existID")
	assert.NotEmpty(t, user)
	assert.NoError(t, err)

	// 異常処理テスト
	// 1. 存在するプロバイダーと存在しないID
	user, err = test.GetByProviderID("google", "notExistID")
	assert.Empty(t, user)
	assert.NoError(t, err)

	// 2. 存在しないプロバイダー
	user, err = test.GetByProviderID("foo", "bar")
	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestCreateUser(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	testDate := time.Date(2000, time.April, 1, 1, 00, 00, 00, jst)
	testutil.SetFakeUuID("testUUID")
	testutil.SetFakeTime(testDate)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user, _ := models.NewUser("testEmail")
	userDetail, _ := models.NewUserDetail("testUUID", "testName")

	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	user.GoogleID = sql.NullString{String: "notExistID", Valid: true}
	mockUserRepository.EXPECT().Insert(user).Return(nil)

	mockUserDetailRepository := mock_repository.NewMockUserDetailRepository(ctrl)
	mockUserDetailRepository.EXPECT().Insert(userDetail).Return(nil)

	mockUserService := mock_service.NewMockIUserService(ctrl)
	mockUserService.EXPECT().IsExist("google", "notExistID").Return(true, nil)
	mockUserService.EXPECT().IsExist("google", "existID").Return(false, nil)

	mockFavoriteGameRepository := mock_repository.NewMockFavoriteGameRepository(ctrl)

	test := NewUserUsecase(mockUserRepository, mockUserDetailRepository, mockUserService, mockFavoriteGameRepository)

	// 正常処理テスト
	body := input.UserCreateReqBody{Provider: "google", ProviderID: "notExistID", UserName: "testName", Email: "testEmail"}
	req, _ := http.NewRequest("GET", "", nil)
	context := &gin.Context{Request: req}
	user, err := test.Create(context, body)

	assert.NotEmpty(t, user)
	assert.NoError(t, err)

	// 異常処理テスト
	// 1. 予想されていないプロバイダー
	body = input.UserCreateReqBody{Provider: "foo", ProviderID: "notExistID", UserName: "testName", Email: "testEmail"}
	req, _ = http.NewRequest("GET", "", nil)
	context = &gin.Context{Request: req}
	user, err = test.Create(context, body)

	assert.Empty(t, user)
	assert.Error(t, err)

	// 2. 存在するIDで作成
	body = input.UserCreateReqBody{Provider: "google", ProviderID: "existID", UserName: "testName", Email: "testEmail"}
	req, _ = http.NewRequest("GET", "", nil)
	context = &gin.Context{Request: req}
	user, err = test.Create(context, body)

	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindByID("existID").Return(&models.User{}, nil)
	mockUserRepository.EXPECT().FindByID("notExistID").Return(nil, nil)

	mockUserDetailRepository := mock_repository.NewMockUserDetailRepository(ctrl)
	mockUserDetailRepository.EXPECT().Update("existID", "testName", "testIcon").Return(nil)

	mockUserService := mock_service.NewMockIUserService(ctrl)
	mockFavoriteGameRepository := mock_repository.NewMockFavoriteGameRepository(ctrl)
	mockFavoriteGameRepository.EXPECT().FindByID("existID").Return(nil, nil)

	test := NewUserUsecase(mockUserRepository, mockUserDetailRepository, mockUserService, mockFavoriteGameRepository)

	existToken := auth.GenerateAccessToken("existID", false)
	notExistToken := auth.GenerateAccessToken("notExistID", false)

	// 正常処理テスト
	bodyReader := strings.NewReader(
		`{
			"user_name": "testName",
			"icon": "testIcon",
			"favorite_games":[]
		}`)
	req, _ := http.NewRequest("GET", "", bodyReader)
	req.Header.Add("Authorization", existToken)
	context := &gin.Context{Request: req}
	err := test.Update(context)

	assert.NoError(t, err)

	// 異常処理テスト
	// 1. 存在しないID
	bodyReader = strings.NewReader(
		`{
			"user_name": "testName",
			"icon": "testIcon",
			"favorite_games":[]
		}`)
	req, _ = http.NewRequest("GET", "", bodyReader)
	req.Header.Add("Authorization", notExistToken)
	context = &gin.Context{Request: req}
	err = test.Update(context)

	assert.Error(t, err)

	// 2. トークン無し
	bodyReader = strings.NewReader(
		`{
			"user_name": "testName",
			"icon": "testIcon",
			"favorite_games":[]
		}`)
	req, _ = http.NewRequest("GET", "", bodyReader)
	context = &gin.Context{Request: req}
	err = test.Update(context)

	assert.Error(t, err)
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindByID("existID").Return(&models.User{}, nil)
	mockUserRepository.EXPECT().FindByID("notExistID").Return(nil, nil)
	mockUserRepository.EXPECT().Delete("existID").Return(nil)

	mockUserDetailRepository := mock_repository.NewMockUserDetailRepository(ctrl)
	mockUserService := mock_service.NewMockIUserService(ctrl)
	mockFavoriteGameRepository := mock_repository.NewMockFavoriteGameRepository(ctrl)

	test := NewUserUsecase(mockUserRepository, mockUserDetailRepository, mockUserService, mockFavoriteGameRepository)

	existToken := auth.GenerateAccessToken("existID", false)
	notExistToken := auth.GenerateAccessToken("notExistID", false)

	// 正常処理テスト
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	context := &gin.Context{Request: req}
	err := test.Delete(context)

	assert.NoError(t, err)

	// 異常処理テスト
	// 1. 存在しないID
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", notExistToken)
	context = &gin.Context{Request: req}
	err = test.Delete(context)

	assert.Error(t, err)

	// 2. トークン無し
	req, _ = http.NewRequest("GET", "", nil)
	context = &gin.Context{Request: req}
	err = test.Delete(context)

	assert.Error(t, err)
}
