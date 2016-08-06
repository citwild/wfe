package mongostore

import (
	"github.com/citwild/wfe/store"
	"gopkg.in/mgo.v2"
)

func NewStores(s *mgo.Session) store.Stores {
	return store.Stores{
		Accounts: NewAccountsStore(s),
		Password: NewPasswordStore(s),
	}
}
