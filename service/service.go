package service

import (
	"github.com/citwild/wfe/api"
)

func NewServers() api.Servers {
	return api.Servers{
		Accounts: NewAccountsServer(),
	}
}
