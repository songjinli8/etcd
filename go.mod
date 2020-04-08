module github.com/songjinli8/etcd

go 1.14

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4

replace github.com/coreos/etcd@v3.3.20+incompatible => github.com/coreos/etcd v3.4.1+incompatible

replace github.com/etcd-io/etcd@v3.3.20+incompatible => github.com/etcd-io/etcd v3.4.1+incompatible

require (
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/etcd-io/etcd v3.3.20+incompatible
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	go.uber.org/zap v1.14.1 // indirect
	google.golang.org/grpc v1.26.0
)
