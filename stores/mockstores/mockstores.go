package mockstores

import (
	"github.com/citwild/wfe/stores"
	"github.com/golang/mock/gomock"
)

func NewStores(ctrl *gomock.Controller) stores.Stores {
	return stores.Stores{
		Accounts: NewMockAccountsStore(ctrl),
		Password: NewMockPasswordStore(ctrl),
	}
}
