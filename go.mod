module li17server

go 1.15

require (
	github.com/ethereum/go-ethereum v1.12.0
	github.com/go-resty/resty/v2 v2.7.0
	github.com/gogf/gf/contrib/drivers/pgsql/v2 v2.5.2
	github.com/gogf/gf/contrib/nosql/redis/v2 v2.5.2
	github.com/gogf/gf/contrib/registry/etcd/v2 v2.5.2
	github.com/gogf/gf/contrib/rpc/grpcx/v2 v2.5.2
	github.com/gogf/gf/contrib/trace/jaeger/v2 v2.5.2
	github.com/gogf/gf/v2 v2.5.2
	github.com/golang/protobuf v1.5.3
	github.com/panjf2000/ants/v2 v2.8.1
	github.com/yitter/idgenerator-go v1.3.3
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

replace github.com/franklihub/mpcCommon v0.0.0 => ./mpcCommon
