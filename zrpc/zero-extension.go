package zrpc

import (
	"github.com/zeromicro/go-zero/zrpc/internal"
)

func WithZero(l internal.ListenFn) internal.ServerOption {
	return internal.WithZero(l)
}
