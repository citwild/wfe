package mockstore

import (
	"github.com/citwild/wfe/store"
	"github.com/golang/mock/gomock"
)

func NewStores(ctrl *gomock.Controller) store.Stores {
	return store.Stores{
		Accounts: NewMockAccountsStore(ctrl),
		Password: NewMockPasswordStore(ctrl),
	}
}
