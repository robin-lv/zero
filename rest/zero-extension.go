package rest

import "github.com/zeromicro/go-zero/rest/internal"

type ListenFn internal.ListenFn

func WithZero(l ListenFn) RunOption {
	return func(server *Server) {
		server.ngin.listenFn = internal.ListenFn(l)
	}
}
