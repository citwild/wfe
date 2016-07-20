package api

import "google.golang.org/grpc"

type Client struct {
	Accounts AccountsClient
}

func NewClient(cc *grpc.ClientConn) *Client {
	c := new(Client)

	c.Accounts = NewAccountsClient(cc)

	return c
}
