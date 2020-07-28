//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// IUserUseCase : インターフェース
type IUserUseCase interface {
	GetMe(c *gin.Context) (*output.UserResBody, error)
	GetByID(c *gin.Context) (*output.UserResBody, error)
	GetByProviderID(provider, providerID string) (*models.User, error)
	Create(c *gin.Context, bpdy input.UserCreateReqBody) (*models.User, error)
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}

type userUsecase struct {
	userRepository         repository.IUserRepository
	userDetailRepository   repository.IUserDetailRepository
	userService            service.IUserService
	favoriteGameRepository repository.IFavoriteGameRepository
}

// NewUserUsecase : Userユースケースの生成
func NewUserUsecase(
	uR repository.IUserRepository,
	uDR repository.IUserDetailRepository,
	uS service.IUserService,
	fGR repository.IFavoriteGameRepository) IUserUseCase {
	return &userUsecase{
		userRepository:         uR,
		userDetailRepository:   uDR,
		userService:            uS,
		favoriteGameRepository: fGR,
	}
}

func (u *userUsecase) GetMe(c *gin.Context) (*output.UserResBody, error) {
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return nil, errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	claimsID := claims["sub"].(string)
	user, err := u.userRepository.FindByID(claimsID)
	if err != nil {
		return nil, err
	}
	ok, err := u.userService.IsDelete(claimsID)
	if !ok {
		return nil, errors.ErrNotFound{}
	}
	userDetail, err := u.userDetailRepository.FindByID(claimsID)
	if err != nil {
		return nil, err
	}
	favoriteGames, err := u.favoriteGameRepository.FindByID(claimsID)
	if err != nil {
		return nil, err
	}
	fmt.Println(favoriteGames)
	resBody := &output.UserResBody{
		UserID:    user.UserID,
		UserName:  userDetail.UserName,
		Icon:      userDetail.Icon,
		CreatedAt: user.CreatedAt,
	}
	for _, g := range favoriteGames {
		r := output.FavoriteGamesRes{
			GameListID: g.GameListID,
			GameTitle:  g.GameTitle,
			CreatedAt:  g.CreatedAt,
		}
		resBody.FavoriteGames = append(resBody.FavoriteGames, r)
	}
	return resBody, nil
}

func (u *userUsecase) GetByID(c *gin.Context) (*output.UserResBody, error) {
	userID := c.Params.ByName("id")
	ok, err := u.userService.IsDelete(userID)
	if !ok {
		return nil, errors.ErrNotFound{}
	}
	user, err := u.userRepository.FindByID(userID)
	if err != nil {
		return nil, err
	}
	userDetail, err := u.userDetailRepository.FindByID(userID)
	if err != nil {
		return nil, err
	}
	favorateGames, err := u.favoriteGameRepository.FindByID(userID)
	if err != nil {
		return nil, err
	}
	resBody := &output.UserResBody{
		UserID:   user.UserID,
		UserName: userDetail.UserName,
		Icon:     userDetail.Icon,
	}
	for _, g := range favorateGames {
		r := output.FavoriteGamesRes{
			GameListID: g.GameTitle,
			GameTitle:  g.GameTitle,
			CreatedAt:  g.CreatedAt,
		}
		resBody.FavoriteGames = append(resBody.FavoriteGames, r)
	}
	return resBody, nil
}

func (u *userUsecase) GetByProviderID(provider, providerID string) (*models.User, error) {
	user, err := u.userRepository.FindByProviderID(provider, providerID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Create(c *gin.Context, body input.UserCreateReqBody) (*models.User, error) {
	user, err := models.NewUser(body.Email)
	if err != nil {
		return nil, err
	}
	switch body.Provider {
	case "google":
		ok, err := u.userService.IsExist(body.Provider, body.ProviderID)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.ErrIDAlreadyExists{}
		}
		user.GoogleID = sql.NullString{String: body.ProviderID, Valid: true}
	default:
		return nil, errors.ErrUnexpectedQueryProvider{Provider: body.Provider}
	}
	userDetail, err := models.NewUserDetail(user.UserID, body.UserName)
	if err != nil {
		return nil, err
	}
	if err := u.userRepository.Insert(user); err != nil {
		return nil, err
	}
	if err := u.userDetailRepository.Insert(userDetail); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Update(c *gin.Context) error {
	body := input.UserUpdateReqBody{}
	if err := c.BindJSON(&body); err != nil {
		return errors.ErrUserUpdateReqBinding{UserName: body.UserName, Icon: body.Icon, FavoriteGames: body.FavoriteGames}
	}
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	claimsID := claims["sub"].(string)
	user, _ := u.userRepository.FindByID(claimsID)
	if user == nil {
		return errors.ErrNotFound{}
	}
	if err := u.userDetailRepository.Update(claimsID, body.UserName, body.Icon); err != nil {
		return err
	}
	favoriteGames, err := u.favoriteGameRepository.FindByID(claimsID)
	if err != nil {
		return err
	}
	beforeFavoriteGames := []input.FavoriteGameRecord{}
	for _, v := range favoriteGames {
		beforeFavoriteGames = append(beforeFavoriteGames, input.FavoriteGameRecord{GameTitle: v.GameTitle})
	}
	afterFavoriteGames := body.FavoriteGames
	var insertGames []input.FavoriteGameRecord
	var deleteGames []input.FavoriteGameRecord
	for _, g := range afterFavoriteGames {
		if !containsRecord(beforeFavoriteGames, g.GameTitle) {
			insertGames = append(insertGames, g)
		}
	}
	for _, g := range beforeFavoriteGames {
		if !containsRecord(afterFavoriteGames, g.GameTitle) {
			deleteGames = append(deleteGames, g)
		}
	}
	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for _, g := range insertGames {
		f, err := models.NewFavoriteGame(claimsID, g.GameTitle)
		if err != nil {
			return err
		}
		eg.Go(func() error {
			if err := u.favoriteGameRepository.Insert(f); err != nil {
				return err
			}
			return nil
		})
	}
	for _, g := range deleteGames {
		g := g
		eg.Go(func() error {
			if err := u.favoriteGameRepository.Delete(claimsID, g.GameTitle); err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		cancel()
		return err
	}
	return nil
}

func (u *userUsecase) Delete(c *gin.Context) error {
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return err
	}
	claimsID := claims["sub"].(string)
	user, err := u.userRepository.FindByID(claimsID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.ErrNotFound{}
	}
	if err := u.userRepository.Delete(claimsID); err != nil {
		return err
	}
	return nil
}
