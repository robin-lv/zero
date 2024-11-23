package internal

import (
	"net"
)

type ListenFn func(network, address string) (net.Listener, error)

type srvOptionExt struct {
	listener net.Listener
}

// WithListener 是一个用于配置服务器监听器
func WithListener(l net.Listener) ServerOption {
	return func(options *rpcServerOptions) { options.ext.listener = l }
}
func (s *baseRpcServer) newListener() (l net.Listener, err error) {
	if s.ext.listener == nil {
		return net.Listen("tcp", s.address)
	}
	l, s.ext.listener = s.ext.listener, nil
	return
}
