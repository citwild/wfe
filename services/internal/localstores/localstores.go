package localstores

import "github.com/citwild/wfe/stores"

func NewLocalStores() stores.Stores {
	return stores.Stores{
		Accounts: &accounts{},
		Password: &password{},
	}
}
