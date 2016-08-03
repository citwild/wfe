package mongostore

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)

type key int

const (
	sessionKey key = iota
)

func WithSession(ctx context.Context, s *mgo.Session) context.Context {
	return context.WithValue(ctx, sessionKey, s)
}

func Session(ctx context.Context) *mgo.Session {
	s, ok := ctx.Value(sessionKey).(*mgo.Session)
	if !ok || s == nil {
		panic("no Session set in context")
	}
	return s
}
