package backend

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/citwild/wfe/api/wfe"
	"github.com/citwild/wfe/pkg/store/mockstore"
	"golang.org/x/net/context"
	"github.com/citwild/wfe/pkg/store"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockAccounts := mockstore.NewMockAccounts(ctrl)
	ctx = store.WithAccounts(ctx, mockAccounts)

	login := "user"
	password := "pass"
	email := "mail@me.com"

	mockAccounts.EXPECT().Create(ctx, &wfe.User{Login: login}, &wfe.EmailAddr{Email: email}).Return(&wfe.User{UID: 123}, nil)

	Accounts.Create(ctx, &wfe.NewAccount{Login: login, Password: password, Email: email})
}
