package api

import (
	"github.com/gorilla/mux"
	"github.com/wide-field-ethnography/wfe/router"
)

func Handler() *mux.Router  {
	m := router.API();
	return m
}