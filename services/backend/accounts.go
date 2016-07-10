package backend

import (
	"golang.org/x/net/context"
	"github.com/wide-field-ethnography/wfe/api/wfe"
)

type accounts struct{}

func (s *accounts) Create(ctx context.Context, newAcct *wfe.NewAccount) (*wfe.CreatedAccount, error) {
	return nil, nil;
}