package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/errors"

	log "github.com/sirupsen/logrus"
)

// IReportHandler : インターフェース
type IReportHandler interface {
	Create(*gin.Context)
}

type reportHandler struct {
	reportUsecase usecase.IReportUsecase
}

// NewReportHanlder :
func NewReportHanlder(rU usecase.IReportUsecase) IReportHandler {
	return &reportHandler{
		reportUsecase: rU,
	}
}

func (rH *reportHandler) Create(c *gin.Context) {
	err := rH.reportUsecase.Create(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrReportReqBinding:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrGenerateID:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success report"})
}
