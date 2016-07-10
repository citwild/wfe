package api

import (
	"github.com/gorilla/mux"
	"citw.uw.edu/wfe/router"
)

func Handler() *mux.Router  {
	m := router.API();
	return m
}