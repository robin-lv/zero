package zrpc

import (
	"net"

	"github.com/zeromicro/go-zero/zrpc/internal"
)

func WithListener(l net.Listener) internal.ServerOption {
	return internal.WithListener(l)
}
