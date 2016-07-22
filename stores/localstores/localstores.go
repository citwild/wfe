package localstores

import "github.com/citwild/wfe/stores"

func NewStores() stores.Stores {
	return stores.Stores{
		Accounts: NewAccountsStore(),
		Password: NewPasswordStore(),
	}
}
