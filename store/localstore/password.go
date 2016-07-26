package localstore

import (
	"errors"
	"github.com/citwild/wfe/store"
	"golang.org/x/net/context"
)

type PasswordStore struct{}

var _ store.PasswordStore = (*PasswordStore)(nil)

func NewPasswordStore() *PasswordStore {
	return &PasswordStore{}
}

func (s PasswordStore) SetPassword(ctx context.Context, UID int32, password string) error {
	return errors.New("Not yet implemented")
}
