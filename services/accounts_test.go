package services

import (
	"testing"

	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/stores"
	"github.com/citwild/wfe/stores/mock_stores"
	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := newTestContext(ctrl)

	accounts := stores.AccountsFromContext(ctx).(*mock_stores.MockAccounts)
	accounts.EXPECT().Create(ctx, &api.User{Login: "me"}, &api.EmailAddr{Email: "e@mail.com"}).
		Return(&api.User{UID: 123}, nil)

	password := stores.PasswordFromContext(ctx).(*mock_stores.MockPassword)
	password.EXPECT().SetPassword(ctx, int32(123), "pass").Return(nil)

	Accounts.Create(ctx, &api.NewAccount{Login: "me", Password: "pass", Email: "e@mail.com"})
}
