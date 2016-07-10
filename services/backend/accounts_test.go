package backend

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/wide-field-ethnography/wfe/api/wfe"
	"github.com/wide-field-ethnography/wfe/pkg/store/mockstore"
	"golang.org/x/net/context"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockAccounts := mockstore.NewMockAccounts(ctrl)
	mockAccounts.EXPECT().Create(ctx, nil, nil).Return(&wfe.User{UID: 123}, nil)

	Accounts.Create(ctx, &wfe.NewAccount{Login: "user", Password: "pass", Email: "email@email.com"})
}
