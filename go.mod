module mpcServer

go 1.20

require (
	github.com/ethereum/go-ethereum v1.13.14
	github.com/gogf/gf/contrib/drivers/pgsql/v2 v2.7.1
	github.com/gogf/gf/contrib/nosql/redis/v2 v2.7.1
	github.com/gogf/gf/contrib/rpc/grpcx/v2 v2.7.1
	github.com/gogf/gf/contrib/trace/jaeger/v2 v2.7.1
	github.com/gogf/gf/v2 v2.7.1
	github.com/golang/protobuf v1.5.3
	github.com/mpcsdk/mpcCommon v0.0.0
	github.com/nats-io/nats.go v1.33.1
	github.com/nats-rpc/nrpc v0.0.0-20231018091755-18e69326f052
	github.com/yitter/idgenerator-go v1.3.3
	go.opentelemetry.io/otel/trace v1.14.0
	google.golang.org/grpc v1.57.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/bits-and-blooms/bitset v1.10.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.12.1 // indirect
	github.com/crate-crypto/go-kzg-4844 v0.7.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/ethereum/c-kzg-4844 v0.4.0 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-resty/resty/v2 v2.11.0 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/gogf/gf/contrib/drivers/mysql/v2 v2.7.1 // indirect
	github.com/gogf/gf/contrib/registry/file/v2 v2.7.1 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/grokify/html-strip-tags-go v0.1.0 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/redis/go-redis/v9 v9.2.1 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/supranational/blst v0.3.11 // indirect
	go.opentelemetry.io/otel v1.14.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.14.0 // indirect
	go.opentelemetry.io/otel/sdk v1.14.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)

replace github.com/mpcsdk/mpcCommon v0.0.0 => ./mpcCommon
