package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	mock "github.com/taniwhy/mochi-match-rest/domain/repository/mock_repository"
)

func TestFindUserByProviderID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := models.User{}

	testMock := mock.NewMockUserRepository(ctrl)
	testMock.EXPECT().FindUserByProviderID("foo", "bar").Return(&u, nil)
	testMock.EXPECT().FindUserByProviderID("foo", "").Return(nil, nil)
	testMock.EXPECT().FindUserByProviderID("", "bar").Return(nil, nil)
	testMock.EXPECT().FindUserByProviderID("", "").Return(nil, nil)

	test := NewUserUsecase(testMock)

	var tests = []struct {
		provider string
		id       string
	}{
		{"foo", "bar"},
		{"foo", ""},
		{"", "bar"},
		{"", ""},
	}
	for _, tt := range tests {
		u, err := test.FindUserByProviderID(tt.provider, tt.id)
		if err != nil {
			t.Fatal(err)
		}
		if tt.provider == "" || tt.id == "" {
			if u != nil {
				t.Errorf("期待する出力: nil, 出力: %v", u)
			}
		}
	}
}
