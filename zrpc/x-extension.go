package zrpc

import (
	"github.com/robin-lv/zero/zrpc/internal"
)

func WithListener(fn internal.ListenFn) internal.ServerOption {
	return internal.WithListener(fn)
}
