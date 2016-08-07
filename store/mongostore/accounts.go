package mongostore

import (
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
	db := newDB(s.session)

	var u dbUser
	u.fromUser(newUser)

	err := db.C("users").Insert(&u)
	if err != nil {
		return nil, err
	}

	// TODO: insert email

	return u.toUser(), nil
}
