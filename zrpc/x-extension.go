package zrpc

import (
	"github.com/zeromicro/go-zero/zrpc/internal"
)

func WithListener(fn internal.ListenFn) internal.ServerOption {
	return internal.WithListener(fn)
}
