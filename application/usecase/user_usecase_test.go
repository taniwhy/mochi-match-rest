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

	// 下記処理の返却値に使用
	u := models.User{}

	// メソッドの行いを定義
	testMock := mock.NewMockUserRepository(ctrl)
	testMock.EXPECT().FindUserByProviderID("foo", "bar").Return(&u, nil)
	testMock.EXPECT().FindUserByProviderID("foo", "").Return(nil, nil)
	testMock.EXPECT().FindUserByProviderID("", "bar").Return(nil, nil)
	testMock.EXPECT().FindUserByProviderID("", "").Return(nil, nil)

	// テスト対象をインスタンス化
	test := NewUserUsecase(testMock)

	// テスト用の引数データの定義
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

func TestCreateUser(t *testing.T) {
	// todo
}
