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
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
	"github.com/taniwhy/mochi-match-rest/util/testutil"

	mock_repository "github.com/taniwhy/mochi-match-rest/domain/repository/mock_repository"
)

func TestCreateReport(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	testDate := time.Date(2020, time.April, 1, 00, 00, 00, 00, jst)
	testutil.SetFakeUuID("testUUID")
	testutil.SetFakeTime(testDate)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newReport, _ := models.NewReport("existID", "VaiolatorID", "existRoomID", "VaiolationDetail")

	mockReportRepository := mock_repository.NewMockIReportRepository(ctrl)
	mockReportRepository.EXPECT().Insert(newReport).Return(nil)

	test := NewReportUsecase(mockReportRepository)

	existToken := auth.GenerateAccessToken("existID", false)

	// 正常処理テスト
	bodyReader := strings.NewReader(`
		{
			"vaiolator_id": "VaiolatorID",
			"detail": "VaiolationDetail"
		}
		`)
	req, _ := http.NewRequest("GET", "", bodyReader)
	req.Header.Add("Authorization", existToken)
	param := gin.Param{Key: "id", Value: "existRoomID"}
	params := gin.Params{param}
	context := &gin.Context{Request: req, Params: params}
	err := test.Create(context)

	assert.NoError(t, err)

	// TODO
	// 異常処理テスト
	// 1. トークン無し

	// 2. 異常なトークン

}
