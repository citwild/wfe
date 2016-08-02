package localstore

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)

type key int

const (
	dbSessionKey key = iota
)

func WithDBSession(ctx context.Context, s *mgo.Session) {
	return context.WithValue(ctx, dbSessionKey, s)
}

func DBSession(ctx context.Context) *mgo.Session {
	s, ok := ctx.Value(dbSessionKey).(*mgo.Session)
	if !ok || s == nil {
		panic("no DB session set in context")
	}
	return s
}
