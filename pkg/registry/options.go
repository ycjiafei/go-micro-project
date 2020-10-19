package registry

import (
	"context"
	"time"
)

type Options struct {
	Addrs []string
	Timeout time.Duration
	Ctx context.Context
}

type RegisterOptions struct {
	TTL time.Duration
	Ctx context.Context
}

type WatchOptions struct {
	Service string
	Ctx context.Context
}

func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}

func RegisterTTL(t time.Duration) RegisterOption {
	return func(o *RegisterOptions) {
		o.TTL = t
	}
}

func WatchService(name string) WatchOption {
	return func(o *WatchOptions) {
		o.Service = name
	}
}