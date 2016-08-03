package mongostore

import "github.com/citwild/wfe/store"

func NewStores() store.Stores {
	return store.Stores{
		Accounts: NewAccountsStore(),
		Password: NewPasswordStore(),
	}
}
