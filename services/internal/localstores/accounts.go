package localstores

import (
	"errors"
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/stores"
	"golang.org/x/net/context"
)

type accounts struct{}

var _ stores.AccountsStore = (*accounts)(nil)

func (s *accounts) Create(ctx context.Context, newUser *api.User, email *api.EmailAddress) (*api.User, error) {
	return nil, errors.New("Not yet implemented")
}
