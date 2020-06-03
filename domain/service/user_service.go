package service

import (
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// IUserService : インターフェース
type IUserService interface {
	IsDelete(id string) (bool, error)
	IsExist(id, provider string) (bool, error)
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService : UserServiceの生成
func NewUserService(uR repository.UserRepository) IUserService {
	return &userService{
		userRepository: uR,
	}
}

func (uS userService) IsDelete(id string) (bool, error) {
	res, err := uS.userRepository.FindByID(id)
	if err != nil {
		return false, err
	}
	if res == nil || res.IsDelete == true {
		return false, nil
	}
	return true, nil
}

func (uS userService) IsExist(id, provider string) (bool, error) {
	res, err := uS.userRepository.FindByProviderID(id, provider)
	if err != nil {
		return false, err
	}
	if res != nil {
		return false, nil
	}
	return true, nil
}
