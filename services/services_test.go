package services

import (
	"github.com/citwild/wfe/stores"
	"github.com/citwild/wfe/stores/mock_stores"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
)

func newTestContext(ctrl *gomock.Controller) context.Context {
	ctx := context.Background()
	ctx = stores.WithAccounts(ctx, mock_stores.NewMockAccounts(ctrl))
	ctx = stores.WithPassword(ctx, mock_stores.NewMockPassword(ctrl))
	return ctx
}
