package mockapi

import (
	"github.com/citwild/wfe/api"
	"github.com/golang/mock/gomock"
)

func NewServers(ctrl *gomock.Controller) api.Servers {
	return api.Servers{
		Accounts: NewMockAccountsServer(ctrl),
	}
}
