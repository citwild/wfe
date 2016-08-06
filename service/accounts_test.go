package service

import (
	"testing"

	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/store"
	"github.com/citwild/wfe/store/mockstore"
	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := testContext(ctrl)

	acct := store.Accounts(ctx).(*mockstore.MockAccountsStore)
	acct.EXPECT().Create(ctx, &api.User{Login: "me"}, &api.EmailAddress{Email: "e@mail.com"}).
		Return(&api.User{UID: 123}, nil)

	pass := store.Password(ctx).(*mockstore.MockPasswordStore)
	pass.EXPECT().SetPassword(ctx, int32(123), "pass")

	actual, err := NewAccountsServer().Create(ctx, &api.NewAccount{Login: "me", Password: "pass", Email: "e@mail.com"})
	if err != nil {
		t.Fatal(err)
	}

	expected := &api.CreatedAccount{UID: 123}
	if *actual != *expected {
		t.Errorf("Account: expected %+v, actual %+v", expected, actual)
	}
}
