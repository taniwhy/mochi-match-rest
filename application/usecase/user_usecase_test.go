package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/taniwhy/mochi-match-rest/domain/repository/mock_repository"
)

func TestFindUserByProviderID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testMock := mock.NewMockUserRepository(ctrl)
	testMock.EXPECT().FindUserByProviderID("google", "1234").Return(nil, nil)

	test := NewUserUsecase(testMock)
	_, err := test.FindUserByProviderID("google", "1234")
	if err != nil {
		t.Error("failed Test")
	}
}
