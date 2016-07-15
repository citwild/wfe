package services

import (
	"github.com/citwild/wfe/stores"
	"github.com/citwild/wfe/stores/mockstores"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
)

func newTestContext(ctrl *gomock.Controller) context.Context {
	ctx := context.Background()
	ctx = stores.WithAccounts(ctx, mockstores.NewMockAccounts(ctrl))
	ctx = stores.WithPassword(ctx, mockstores.NewMockPassword(ctrl))
	return ctx
}
