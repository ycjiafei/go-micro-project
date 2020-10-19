package service

import "github.com/ycjiafei/go-micro-project/pkg/registry"

type serviceRegistry struct {
	opts registry.Options
	name string
	address string
}


