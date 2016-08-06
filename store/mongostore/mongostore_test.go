package mongostore

import (
	"github.com/citwild/wfe/store"
	"github.com/citwild/wfe/store/mockstore"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
)

func testContext(ctrl *gomock.Controller) context.Context {
	ctx := context.Background()

	ctx = store.WithStores(ctx, mockstore.NewStores(ctrl))

	return ctx
}
