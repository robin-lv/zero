package zrpc

import (
	"github.com/zeromicro/go-zero/zrpc/internal"
)

type ListenFn internal.ListenFn

func WithZero(l ListenFn) internal.ServerOption {
	return internal.WithZero(internal.ListenFn(l))
}
