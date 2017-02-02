package app

import (
	"sync"
	"github.com/citwild/wfe/app/internal/tmpl"
)

var initOnce sync.Once

func Init() {
	initOnce.Do(func() {
		tmpl.Load()
	})
}