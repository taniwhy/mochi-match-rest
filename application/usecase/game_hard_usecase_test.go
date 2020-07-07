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

func TestFindAllGameHard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGameHardRepository := mock_repository.NewMockIGameHardRepository(ctrl)
	mockGameHardRepository.EXPECT().FindAll().Return([]*models.GameHard{}, nil)

	test := NewGameHardUsecase(mockGameHardRepository)

	// 正常処理
	req, _ := http.NewRequest("GET", "", nil)
	context := &gin.Context{Request: req}
	gamehards, err := test.FindAll(context)

	assert.NotNil(t, gamehards)
	assert.NoError(t, err)
}

func TestInsertGameHard(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	testDate := time.Date(2020, time.April, 1, 00, 00, 00, 00, jst)
	testutil.SetFakeUuID("testUUID")
	testutil.SetFakeTime(testDate)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newGameList, _ := models.NewGameHard("hardname")

	mockGameHardRepository := mock_repository.NewMockIGameHardRepository(ctrl)
	mockGameHardRepository.EXPECT().Insert(newGameList).Return(nil)

	test := NewGameHardUsecase(mockGameHardRepository)

	// 正常処理
	bodyReader := strings.NewReader(`
		{
			"hard_name": "hardname"
		}
		`)
	req, _ := http.NewRequest("GET", "", bodyReader)
	context := &gin.Context{Request: req}
	err := test.Insert(context)

	assert.NoError(t, err)
}

func TestUpdateGameHard(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	testDate := time.Date(2020, time.April, 1, 00, 00, 00, 00, jst)
	testutil.SetFakeTime(testDate)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	updatedGamehard := &models.GameHard{
		GameHardID: "gamehardID",
		HardName:   "hardname",
		UpdateAt:   clock.Now(),
	}
	mockGameHardRepository := mock_repository.NewMockIGameHardRepository(ctrl)
	mockGameHardRepository.EXPECT().Update(updatedGamehard).Return(nil)

	test := NewGameHardUsecase(mockGameHardRepository)

	// 正常処理
	bodyReader := strings.NewReader(`
		{
			"hard_name": "hardname"
		}
		`)
	req, _ := http.NewRequest("GET", "", bodyReader)
	param := gin.Param{Key: "id", Value: "gamehardID"}
	params := gin.Params{param}
	context := &gin.Context{Request: req, Params: params}
	err := test.Update(context)

	assert.NoError(t, err)

	// 異常処理
	// 1. Param無し
	bodyReader = strings.NewReader(`
		{
			"hard_name": "hardname"
		}
		`)
	req, _ = http.NewRequest("GET", "", bodyReader)
	context = &gin.Context{Request: req}
	err = test.Update(context)

	assert.Error(t, err)
}

func TestDeleteGameHard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGameHardRepository := mock_repository.NewMockIGameHardRepository(ctrl)

	test := NewGameHardUsecase(mockGameHardRepository)

	// TODO 配列上のParamを取得できないエラー
	// 正常処理
	req, _ := http.NewRequest("GET", "", nil)
	param := gin.Param{Key: "id", Value: "gamehardID"}
	params := gin.Params{param, param, param}
	context := &gin.Context{Request: req, Params: params}
	err := test.Delete(context)

	assert.NoError(t, err)
}
