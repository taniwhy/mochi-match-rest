//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package service

import (
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// IUserService : インターフェース
type IUserService interface {
	IsAdmin(id string) (bool, error)
	IsDelete(id string) (bool, error)
	IsExist(id, provider string) (bool, error)
}

type userService struct {
	userRepository repository.IUserRepository
}

// NewUserService : UserServiceの生成
func NewUserService(uR repository.IUserRepository) IUserService {
	return &userService{
		userRepository: uR,
	}
}

func (s *userService) IsAdmin(id string) (bool, error) {
	res, err := s.userRepository.FindByID(id)
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, nil
	}
	return res.IsAdmin, nil
}

func (s *userService) IsDelete(id string) (bool, error) {
	res, err := s.userRepository.FindByID(id)
	if err != nil {
		return false, err
	}
	if res == nil || res.IsDelete == true {
		return false, nil
	}
	return true, nil
}

// ユーザーが存在しなければ真を返却
func (s *userService) IsExist(id, provider string) (bool, error) {
	res, err := s.userRepository.FindByProviderID(id, provider)
	if err != nil {
		return false, err
	}
	if res != nil {
		return false, nil
	}
	return true, nil
}
