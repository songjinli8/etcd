module github.com/songjinli8/etcd

go 1.14

replace github.com/etcd-io/etcd => github.com/fredczj/etcd v3.4.6+incompatible

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4

require (
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/coreos/go-systemd/v22 v22.0.0 // indirect
	github.com/etcd-io/etcd v0.0.0-00010101000000-000000000000
	github.com/fredczj/etcd v3.4.7+incompatible // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	go.etcd.io/etcd v3.3.20+incompatible // indirect
	go.uber.org/zap v1.14.1 // indirect
	google.golang.org/grpc v1.26.0
)
