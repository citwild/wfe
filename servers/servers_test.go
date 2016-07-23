package servers

import (
	"github.com/citwild/wfe/api/mockapi"
	"github.com/citwild/wfe/stores"
	"github.com/citwild/wfe/stores/mockstores"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
)

func testContext(ctrl *gomock.Controller) context.Context {
	ctx := context.Background()
	ctx = WithServers(ctx, mockapi.NewServers(ctrl))
	ctx = stores.WithStores(ctx, mockstores.NewStores(ctrl))
	return ctx
}
