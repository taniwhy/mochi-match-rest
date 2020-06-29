//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// IReportUsecase : インターフェース
type IReportUsecase interface {
	Insert(c *gin.Context) error
}

type reportUsecase struct {
	reportRepository repository.IReportRepository
}

// NewReportUsecase : Reportユースケースの生成
func NewReportUsecase(rR repository.IReportRepository) IReportUsecase {
	return &reportUsecase{reportRepository: rR}
}

func (u *reportUsecase) Insert(c *gin.Context) error {
	body := input.ReportReqBody{}
	if err := c.BindJSON(&body); err != nil {
		return errors.ErrReportReqBinding{
			VaiolatorID:      body.VaiolatorID,
			VaiolationDetail: body.VaiolationDetail,
		}
	}
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	callerID := claims["sub"].(string)
	roomID := c.Params.ByName("id")
	report, err := models.NewReport(callerID, body.VaiolatorID, roomID, body.VaiolationDetail)
	if err != nil {
		return err
	}
	if err := u.reportRepository.InsertReport(report); err != nil {
		return err
	}
	return nil
}
