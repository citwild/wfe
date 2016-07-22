package servers

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

	accounts := stores.Accounts(ctx).(*mockstores.MockAccountsStore)
	accounts.EXPECT().Create(ctx, &api.User{Login: "me"}, &api.EmailAddress{Email: "e@mail.com"}).
		Return(&api.User{UID: 123}, nil)

	password := stores.Password(ctx).(*mockstores.MockPasswordStore)
	password.EXPECT().SetPassword(ctx, int32(123), "pass")

	a, err := Accounts(ctx).Create(ctx, &api.NewAccount{Login: "me", Password: "pass", Email: "e@mail.com"})
	if err != nil {
		t.Fatal(err)
	}

	w := &api.CreatedAccount{UID: 123}
	if *a != *w {
		t.Errorf("got %+v, want %+v", a, w)
	}
}
