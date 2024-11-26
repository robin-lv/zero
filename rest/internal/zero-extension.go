package internal

import (
	"net"
	"net/http"
)

type ListenFn func(network, addr string) (net.Listener, error)

func ListenHttp(fn ListenFn, host string, port int, handler http.Handler, opts ...StartOption) error {
	if fn == nil {
		return StartHttp(host, port, handler, opts...)
	}
	return start(host, port, handler, func(svr *http.Server) error {
		l, err := fn("tcp", svr.Addr)
		if err != nil {
			return err
		}
		return svr.Serve(l)
	}, opts...)
}

// ListenHttps starts a https server.
func ListenHttps(fn ListenFn, host string, port int, certFile, keyFile string, handler http.Handler,
	opts ...StartOption) error {
	if fn == nil {
		return StartHttps(host, port, certFile, keyFile, handler, opts...)
	}
	return start(host, port, handler, func(svr *http.Server) error {
		l, err := fn("tcp", svr.Addr)
		if err != nil {
			return err
		}
		// certFile and keyFile are set in buildHttpsServer
		return svr.ServeTLS(l, certFile, keyFile)
	}, opts...)
}
