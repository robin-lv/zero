package internal

import (
	"net"
)

type ListenFn func(network, address string) (net.Listener, error)

func WithZero(l ListenFn) ServerOption {
	return func(options *rpcServerOptions) {
		options.listenFn = l
	}
}
