package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// UserDetailUseCase :
type UserDetailUseCase interface {
	FindUserDetailByID(id string) (*dbmodel.UserDetail, error)
	CreateUserDetail(userDetail *dbmodel.UserDetail) error
	UpdateUserDetail(userDetail *dbmodel.UserDetail) error
	DeleteUserDetail(userDetail *dbmodel.UserDetail) error
}

type userDetailUsecase struct {
	userDetailRepository repository.UserDetailRepository
}

// NewUserDetailUsecase :
func NewUserDetailUsecase(uR repository.UserDetailRepository) UserDetailUseCase {
	return &userDetailUsecase{
		userDetailRepository: uR,
	}
}

func (uU userDetailUsecase) FindUserDetailByID(id string) (*dbmodel.UserDetail, error) {
	userDetail, err := uU.userDetailRepository.FindUserDetailByID(id)
	if err != nil {
		return nil, err
	}
	return userDetail, nil
}

func (uU userDetailUsecase) CreateUserDetail(userDetail *dbmodel.UserDetail) error {
	err := uU.userDetailRepository.InsertUserDetail(userDetail)
	if err != nil {
		return err
	}
	return nil
}

func (uU userDetailUsecase) UpdateUserDetail(userDetail *dbmodel.UserDetail) error {
	err := uU.userDetailRepository.UpdateUserDetail(userDetail)
	if err != nil {
		return err
	}
	return nil
}

func (uU userDetailUsecase) DeleteUserDetail(userDetail *dbmodel.UserDetail) error {
	err := uU.userDetailRepository.InsertUserDetail(userDetail)
	if err != nil {
		return err
	}
	return nil
}
