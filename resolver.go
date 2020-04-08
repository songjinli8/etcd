package etcd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/etcd-io/etcd/clientv3"
	"google.golang.org/grpc/naming"
	//"google.golang.org/grpc/resolver"
)

var cli *clientv3.Client

type resolver struct {
	serviceName string
}

func NewResolver(etcdAddr string) *resolver {
	return &resolver{
		serviceName: etcdAddr,
	}
}

func (r *resolver) Resolve(target string) (naming.Watcher, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(target, ","),
	})
	if err != nil {
		return nil, errors.New("no service name provided")
	}

	return &watcher{re: r, client: *client, isInitialized: false}, nil
}

var Prefix = "etcd3_naming"

type watcher struct {
	re            *resolver
	client        clientv3.Client
	isInitialized bool
}

func (w *watcher) Close() {
}

func (w *watcher) Next() ([]*naming.Update, error) {
	prefix := fmt.Sprintf("%s/%s", Prefix, w.re.serviceName)

	if !w.isInitialized {
		resp, err := w.client.Get(context.Background(), prefix, clientv3.WithPrefix())
		w.isInitialized = true
		if err != nil {
			return nil, errors.New("fail to get prefix")
		}
		addrs := extractAddrs(resp)
		if l := len(addrs); l != 0 {
			updates := make([]*naming.Update, l)
			for i := range addrs {
				updates[i] = &naming.Update{Op: naming.Add, Addr: addrs[i]}
			}

			return updates, nil
		}
	}

	rch := w.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for respEve := range rch {
		for _, ev := range respEve.Events {
			switch ev.Type {
			case mvccpb.PUT:
				return []*naming.Update{{Op: naming.Add, Addr: string(ev.Kv.Value)}}, nil
			case mvccpb.DELETE:
				return []*naming.Update{{Op: naming.Delete, Addr: string(ev.Kv.Value)}}, nil
			}
		}
	}

	return nil, nil
}

func extractAddrs(resp *clientv3.GetResponse) []string {
	addrs := []string{}

	if resp == nil || resp.Kvs == nil {
		return addrs
	}

	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			addrs = append(addrs, string(v))
		}
	}

	return addrs
}
