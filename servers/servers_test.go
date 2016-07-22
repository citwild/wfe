package servers

import (
	"github.com/citwild/wfe/stores"
	"github.com/citwild/wfe/stores/mockstores"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
)

func newTestContext(ctrl *gomock.Controller) context.Context {
	ctx := context.Background()
	ctx = WithServers(ctx, NewServers())
	ctx = stores.WithStores(ctx, mockstores.NewMockStores(ctrl))
	return ctx
}
