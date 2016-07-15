package services

import (
	"testing"

	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/stores"
	"github.com/citwild/wfe/stores/mockstores"
	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := newTestContext(ctrl)

	accounts := stores.AccountsFromContext(ctx).(*mockstores.MockAccounts)
	accounts.EXPECT().Create(ctx, &api.User{Login: "me"}, &api.EmailAddress{Email: "e@mail.com"}).
		Return(&api.User{UID: 123}, nil)

	password := stores.PasswordFromContext(ctx).(*mockstores.MockPassword)
	password.EXPECT().SetPassword(ctx, int32(123), "pass")

	Accounts.Create(ctx, &api.NewAccount{Login: "me", Password: "pass", Email: "e@mail.com"})
}
