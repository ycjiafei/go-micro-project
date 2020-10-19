package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"github.com/ycjiafei/go-micro-project/pkg/registry"
	"sync"
	"time"
)

type etcdRegistry struct {
	client *clientv3.Client
	options registry.Options
	register map[string]*registry.Node  // 保存在上面注册的服务

	sync.RWMutex
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	// 默认值
	e := &etcdRegistry{
		options: registry.Options{
			Timeout: 5 * time.Second,
			Ctx: context.Background(),
			Addrs: []string{"127.0.0.1:2379"},
		},
		register: make(map[string]*registry.Node),
	}
	// 赋值
	for _, o := range opts {
		o(&e.options)
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: e.options.Addrs,
	})
	if err != nil {
		return nil
	}
	e.client = cli
	return e
}

func (e *etcdRegistry) Register(s *registry.Service, opts ...registry.RegisterOption) error {
	if len(s.Nodes) == 0 {
		return errors.New("服务至少包含一个节点")
	}
	for _, n := range s.Nodes {
		e.RLock()
		node, ok := e.register[s.Name + n.Id]
		e.RUnlock()
		if ok {
			n = node
		}
		if err := e.registryEtcd(s, n); err != nil {
			return err
		}
	}

	return nil
}

func (e *etcdRegistry) registryEtcd(s *registry.Service, node *registry.Node) error {
	// 没有注册过
	ctx, cancel := context.WithTimeout(context.Background(), e.options.Timeout)
	defer cancel()
	n, _ := json.Marshal(node)
	if _, err := e.client.Put(ctx, s.Name+node.Id, string(n)); err != nil {
		return err
	}

	e.Lock()
	e.register[s.Name+node.Id] = node
	e.Unlock()

	return nil
}


func decodeService(ds []byte) *registry.Service {
	s :=  &registry.Service{}
	json.Unmarshal(ds, s)
	return s
}

func (e *etcdRegistry) Deregister(registryName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), e.options.Timeout)
	defer cancel()
	e.Lock()
	delete(e.register, registryName)
	e.Unlock()
	_, err := e.client.Delete(ctx, registryName)
	if err != nil {
		return err
	}

	return nil
}