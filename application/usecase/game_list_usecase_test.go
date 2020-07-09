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
	"github.com/taniwhy/mochi-match-rest/domain/repository/mock_repository"
	"github.com/taniwhy/mochi-match-rest/util/clock"
	"github.com/taniwhy/mochi-match-rest/util/testutil"
)

func TestFindAllGameList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGameListRepository := mock_repository.NewMockIGameListRepository(ctrl)
	mockGameListRepository.EXPECT().FindAll().Return([]*models.GameList{}, nil)

	test := NewGameListUsecase(mockGameListRepository)

	// 正常処理
	req, _ := http.NewRequest("GET", "", nil)
	context := &gin.Context{Request: req}
	gamelists, err := test.FindAll(context)

	assert.NotNil(t, gamelists)
	assert.NoError(t, err)
}

func TestInsertGameList(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	testDate := time.Date(2020, time.April, 1, 00, 00, 00, 00, jst)
	testutil.SetFakeUuID("testUUID")
	testutil.SetFakeTime(testDate)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newGameList, _ := models.NewGameList("gametitle")

	mockGameListRepository := mock_repository.NewMockIGameListRepository(ctrl)
	mockGameListRepository.EXPECT().Insert(newGameList).Return(nil)

	test := NewGameListUsecase(mockGameListRepository)

	// 正常処理
	bodyReader := strings.NewReader(`
		{
			"game_title": "gametitle"
		}
		`)
	req, _ := http.NewRequest("GET", "", bodyReader)
	context := &gin.Context{Request: req}
	err := test.Insert(context)

	assert.NoError(t, err)
}

func TestUpdateGameList(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	testDate := time.Date(2020, time.April, 1, 00, 00, 00, 00, jst)
	testutil.SetFakeTime(testDate)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	updatedGamelist := &models.GameList{
		GameListID: "gamelistID",
		GameTitle:  "gametitle",
		UpdateAt:   clock.Now(),
	}
	mockGameListRepository := mock_repository.NewMockIGameListRepository(ctrl)
	mockGameListRepository.EXPECT().Update(updatedGamelist).Return(nil)

	test := NewGameListUsecase(mockGameListRepository)

	// 正常処理
	bodyReader := strings.NewReader(`
		{
			"game_title": "gametitle"
		}
		`)
	req, _ := http.NewRequest("GET", "", bodyReader)
	param := gin.Param{Key: "id", Value: "gamelistID"}
	params := gin.Params{param}
	context := &gin.Context{Request: req, Params: params}
	err := test.Update(context)

	assert.NoError(t, err)

	// 異常処理
	// 1. Param無し
	bodyReader = strings.NewReader(`
		{
			"game_title": "gametitle"
		}
		`)
	req, _ = http.NewRequest("GET", "", bodyReader)
	context = &gin.Context{Request: req}
	err = test.Update(context)

	assert.Error(t, err)
}

func TestDeleteGameList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGameListRepository := mock_repository.NewMockIGameListRepository(ctrl)

	test := NewGameListUsecase(mockGameListRepository)

	// TODO 配列上のParamを取得できないエラー
	// 正常処理
	req, _ := http.NewRequest("GET", "", nil)
	param := gin.Param{Key: "id", Value: "gamelistID"}
	params := gin.Params{param, param, param}
	context := &gin.Context{Request: req, Params: params}
	err := test.Delete(context)

	assert.NoError(t, err)
}
