package rest

import "github.com/zeromicro/go-zero/rest/internal"

func WithZero(l internal.ListenFn) RunOption {
	return func(server *Server) {
		server.ngin.listenFn = l
	}
}
