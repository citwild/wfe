package mongostore

import (
	"errors"
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/store"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)

type AccountsStore struct {
	session *mgo.Session
}

var _ store.AccountsStore = (*AccountsStore)(nil)

func NewAccountsStore(s *mgo.Session) *AccountsStore {
	return &AccountsStore{session: s}
}

func (s *AccountsStore) Create(ctx context.Context, newUser *api.User, email *api.EmailAddress) (*api.User, error) {
	return nil, errors.New("Not yet implemented")
}
