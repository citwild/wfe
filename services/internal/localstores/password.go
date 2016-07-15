package localstores

import (
	"errors"
	"github.com/citwild/wfe/stores"
	"golang.org/x/net/context"
)

type password struct{}

var _ stores.Password = (*password)(nil)

func SetPassword(ctx context.Context, UID int32, password string) error {
	return errors.New("Not yet implemented")
}
