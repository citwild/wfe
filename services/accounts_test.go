package services

import (
	"testing"

	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/stores"
	"github.com/citwild/wfe/stores/mock_stores"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockAccounts := mock_stores.NewMockAccounts(ctrl)
	ctx = stores.WithAccounts(ctx, mockAccounts)

	login := "user"
	password := "pass"
	email := "mail@me.com"

	mockAccounts.EXPECT().Create(ctx, &api.User{Login: login}, &api.EmailAddr{Email: email}).Return(&api.User{UID: 123}, nil)

	Accounts.Create(ctx, &api.NewAccount{Login: login, Password: password, Email: email})
}
