package usecase

import "github.com/taniwhy/mochi-match-rest/domain/repository"

type UserUseCaseInterface interface {
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(uR repository.UserRepository) UserUseCaseInterface {
	return &userUsecase{
		userRepository: uR,
	}
}
