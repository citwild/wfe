package backend

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/citwild/wfe/api"
	"golang.org/x/net/context"
	"github.com/citwild/wfe/store/mock_store"
	"github.com/citwild/wfe/store"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockAccounts := mock_store.NewMockAccounts(ctrl)
	ctx = store.WithAccounts(ctx, mockAccounts)

	login := "user"
	password := "pass"
	email := "mail@me.com"

	mockAccounts.EXPECT().Create(ctx, &api.User{Login: login}, &api.EmailAddr{Email: email}).Return(&api.User{UID: 123}, nil)

	Accounts.Create(ctx, &api.NewAccount{Login: login, Password: password, Email: email})
}
