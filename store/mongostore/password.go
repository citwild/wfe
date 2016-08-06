package mongostore

import (
	"errors"
	"github.com/citwild/wfe/store"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)

type PasswordStore struct {
	session *mgo.Session
}

var _ store.PasswordStore = (*PasswordStore)(nil)

func NewPasswordStore(s *mgo.Session) *PasswordStore {
	return &PasswordStore{session: s}
}

func (s PasswordStore) SetPassword(ctx context.Context, UID int32, password string) error {
	return errors.New("Not yet implemented")
}
