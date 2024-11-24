package internal

import (
	"net"
)

type ListenFn func(network, address string) (net.Listener, error)

type srvOptionExt struct {
	listenFn ListenFn
}

// WithListener 是一个用于配置服务器监听器
func WithListener(fn ListenFn) ServerOption {
	return func(options *rpcServerOptions) { options.ext.listenFn = fn }
}
func (s *baseRpcServer) newListener() (l net.Listener, err error) {
	fn := net.Listen
	if s.ext.listenFn != nil {
		fn, s.ext.listenFn = s.ext.listenFn, nil
	}
	return fn("tcp", s.address)
}
