package internal

import (
	"net/http"
)

var Handlers = map[string]http.Handler{}
