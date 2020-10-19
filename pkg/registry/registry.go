package registry

type Registry interface {
	Register(*Service, ...RegisterOption) error
	Deregister(string) error
}

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

type WatchOption func(*WatchOptions)