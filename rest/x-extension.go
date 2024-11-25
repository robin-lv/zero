package rest

import (
	"github.com/robin-lv/zero/rest/internal"
)

func WithListenFn(fn internal.ListenFn) RunOption {
	return func(server *Server) {
		server.ngin.listenFn = fn
	}
}
