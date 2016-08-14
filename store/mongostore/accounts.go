package mongostore

import (
	"errors"
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/store"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AccountsStore struct {
	session *mgo.Session
}

var _ store.AccountsStore = (*AccountsStore)(nil)

func NewAccountsStore(s *mgo.Session) *AccountsStore {
	return &AccountsStore{session: s}
}

func (s *AccountsStore) Create(ctx context.Context, newUser *api.User, email *api.EmailAddress) (*api.User, error) {
	if newUser.ID != "" {
		return nil, errors.New("ID cannot be set")
	}
	if newUser.Login == "" {
		return nil, errors.New("Login must be set")
	}

	db := newDB(s.session)

	var du dbUser
	err := du.fromUser(newUser)
	if err != nil {
		return nil, err
	}

	du.ID = bson.NewObjectId()
	err = db.C("users").Insert(&du)
	if err != nil {
		return nil, err
	}

	// TODO: insert email

	return du.toUser()
}
