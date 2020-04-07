package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc/resolver"
)

var cli *clientv3.Client

type etcdResolver struct {
	rawAddr string
	cc      resolver.ClientConn
}

func NewResolver(etcdAddr string) resolver.Builder {
	return &etcdResolver{
		rawAddr: etcdAddr,
	}
}

func (r *etcdResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var err error

	if cli == nil {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:
		})
	}
}

func (r *etcdResolver) Scheme() string {

}
