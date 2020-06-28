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
	Insert(*gin.Context) error
}

type reportUsecase struct {
	reportRepository repository.IReportRepository
}

// NewReportUsecase : Reportユースケースの生成
func NewReportUsecase(rR repository.IReportRepository) IReportUsecase {
	return &reportUsecase{
		reportRepository: rR,
	}
}

func (rU reportUsecase) Insert(c *gin.Context) error {
	b := input.ReportReqBody{}
	if err := c.BindJSON(&b); err != nil {
		return errors.ErrReportReqBinding{
			VaiolatorID:      b.VaiolatorID,
			VaiolationDetail: b.VaiolationDetail,
		}
	}
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	cID := claims["sub"].(string)
	rID := c.Params.ByName("id")
	r, err := models.NewReport(cID, b.VaiolatorID, rID, b.VaiolationDetail)
	if err != nil {
		return err
	}
	if err := rU.reportRepository.InsertReport(r); err != nil {
		return err
	}
	return nil
}
