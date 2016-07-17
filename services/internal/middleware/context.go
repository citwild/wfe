package middleware

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func InitContext(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return nil, errors.New("Not yet implemented")
}
