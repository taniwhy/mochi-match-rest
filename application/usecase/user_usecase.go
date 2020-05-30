package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// UserUseCase :
type UserUseCase interface {
	FindUserByProviderID(provider, id string) (*dbmodel.User, error)
	CreateUser(user *dbmodel.User) error
	DeleteUser(id string) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

// NewUserUsecase :
func NewUserUsecase(uR repository.UserRepository) UserUseCase {
	return &userUsecase{
		userRepository: uR,
	}
}

func (uU userUsecase) FindUserByProviderID(provider, id string) (*dbmodel.User, error) {
	user, err := uU.userRepository.FindUserByProviderID(provider, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uU userUsecase) CreateUser(user *dbmodel.User) error {
	err := uU.userRepository.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (uU userUsecase) DeleteUser(id string) error {
	err := uU.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
