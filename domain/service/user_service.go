//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package service

import (
	"fmt"

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

func (uS userService) IsAdmin(id string) (bool, error) {
	res, err := uS.userRepository.FindByID(id)
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, nil
	}
	fmt.Println("んあああああ", res.IsAdmin)
	return res.IsAdmin, nil
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

// ユーザーが存在しなければ真を返却
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
