package app

import (
	"github.com/citwild/wfe/app/internal/tmpl"
	"sync"
)

var initOnce sync.Once

func Init() {
	initOnce.Do(func() {
		tmpl.Load()
	})
}
