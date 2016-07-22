package services

import (
	"github.com/citwild/wfe/api"
)

type Services struct {
	Accounts api.AccountsServer
}

func NewServices() Services {
	return Services{
		Accounts: &accounts{},
	}
}
