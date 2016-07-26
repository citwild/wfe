package localstore

import (
	"errors"
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/store"
	"golang.org/x/net/context"
)

type AccountsStore struct{}

var _ store.AccountsStore = (*AccountsStore)(nil)

func NewAccountsStore() *AccountsStore {
	return &AccountsStore{}
}

func (s *AccountsStore) Create(ctx context.Context, newUser *api.User, email *api.EmailAddress) (*api.User, error) {
	return nil, errors.New("Not yet implemented")
}
