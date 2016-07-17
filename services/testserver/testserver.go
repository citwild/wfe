package testserver

import (
	"errors"
	"github.com/citwild/wfe/api"
	"golang.org/x/net/context"
)

type TestServer struct {
	Client  *api.Client
	Context context.Context
}

func New() *TestServer {
	return &TestServer{}
}

func (s *TestServer) Start() error {
	return errors.New("Not yet implemented")
}

func (s *TestServer) Close() {

}
