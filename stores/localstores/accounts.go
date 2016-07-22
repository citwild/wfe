package localstores

import (
	"errors"
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/stores"
	"golang.org/x/net/context"
)

type AccountsStore struct{}

var _ stores.AccountsStore = (*AccountsStore)(nil)

func (s *AccountsStore) Create(ctx context.Context, newUser *api.User, email *api.EmailAddress) (*api.User, error) {
	return nil, errors.New("Not yet implemented")
}
