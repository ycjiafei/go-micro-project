package test

import (
	"fmt"
	"github.com/ycjiafei/go-micro-project/pkg/registry"
	"github.com/ycjiafei/go-micro-project/pkg/registry/etcd"
	"testing"
)

func TestRegistry(t *testing.T) {
	srv := registry.Service{
		Name: "user",
		Nodes: []*registry.Node{{
			Id: "1",
			Address: "user1.com",
		}},
	}
	r := etcd.NewRegistry()
	err := r.Register(&srv)
	r.Deregister(srv.Name + "1")
	fmt.Println(err)
}
