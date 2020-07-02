package usecase

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_repository "github.com/taniwhy/mochi-match-rest/domain/repository/mock_repository"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
	"github.com/taniwhy/mochi-match-rest/util/testutil"
)

func TestCreateReport(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	testDate := time.Date(2020, time.April, 1, 00, 00, 00, 00, jst)
	testutil.SetFakeUuID("testUUID")
	testutil.SetFakeTime(testDate)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReportRepository := mock_repository.NewMockIReportRepository(ctrl)

	test := NewReportUsecase(mockReportRepository)

	existToken := auth.GenerateAccessToken("existID", false)

	// 正常処理テスト
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	param := gin.Param{Key: "id", Value: "existRoomID"}
	params := gin.Params{param}
	context := &gin.Context{Request: req, Params: params}
	err := test.Create(context)

	assert.NoError(t, err)

	// 異常処理テスト
	// 1. トークン無し
	req, _ = http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", existToken)
	param = gin.Param{Key: "id", Value: "existRoomID"}
	params = gin.Params{param}
	context = &gin.Context{Request: req, Params: params}
	err = test.Create(context)

	assert.Error(t, err)
}
