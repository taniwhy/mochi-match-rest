package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// IEntryHistoryHandler : インターフェース
type IEntryHistoryHandler interface {
	GetByID(*gin.Context)
}

type entryHistoryHandler struct {
	entryHistoryUsecase usecase.IEntryHistoryUseCase
}

// NewEntryHistoryHandler : 参加履歴ハンドラの生成
func NewEntryHistoryHandler(eU usecase.IEntryHistoryUseCase) IEntryHistoryHandler {
	return &entryHistoryHandler{
		entryHistoryUsecase: eU,
	}
}

func (h *entryHistoryHandler) GetByID(c *gin.Context) {
	histories, err := h.entryHistoryUsecase.GetByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, histories)
}
