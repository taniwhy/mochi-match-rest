package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/errors"

	log "github.com/sirupsen/logrus"
)

// IUserHandler : インターフェース
type IUserHandler interface {
	GetMe(*gin.Context)
	GetByID(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type userHandler struct {
	userUsecase usecase.IUserUseCase
}

// NewUserHandler : ユーザーハンドラの生成
func NewUserHandler(uU usecase.IUserUseCase) IUserHandler {
	return &userHandler{
		userUsecase: uU,
	}
}

func (uH userHandler) GetMe(c *gin.Context) {
	u, err := uH.userUsecase.GetMe(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"code": 1, "message": err.Error()})
			return
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"code": 2, "message": err.Error()})
			return
		case errors.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"code": 3, "message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 4, "message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 99, "message": err.Error()})
			log.Fatal("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, u)
}

func (uH userHandler) GetByID(c *gin.Context) {
	u, err := uH.userUsecase.GetByID(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"code": 1, "message": err.Error()})
			return
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"code": 2, "message": err.Error()})
			return
		case errors.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"code": 3, "message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 4, "message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 99, "message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, u)
}

func (uH userHandler) Update(c *gin.Context) {
	err := uH.userUsecase.Update(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"code": 1, "message": err.Error()})
			return
		case errors.ErrUserUpdateReqBinding:
			c.JSON(http.StatusBadRequest, gin.H{"code": 2, "message": err.Error()})
			return
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"code": 3, "message": err.Error()})
			return
		case errors.ErrParams:
			c.JSON(http.StatusBadRequest, gin.H{"code": 4, "message": err.Error()})
			return
		case errors.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"code": 5, "message": err.Error()})
			return
		case errors.ErrGenerateID:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 6, "message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 7, "message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 99, "message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated user"})
}

func (uH userHandler) Delete(c *gin.Context) {
	err := uH.userUsecase.Delete(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"code": 1, "message": err.Error()})
			return
		case errors.ErrParams:
			c.JSON(http.StatusBadRequest, gin.H{"code": 2, "message": err.Error()})
			return
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"code": 3, "message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 99, "message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted user"})
}
