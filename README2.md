# block chain go example

Based in

- https://github.com/plutov/packagemain
- https://www.youtube.com/watch?v=gju-bml4kdw

## structure

```shell
├── Gopkg.lock
├── Gopkg.toml # dependencies
├── blockchain
│   └── blockchain.go
├── client
│   └── main.go
├── debug
├── main.go
├── metrics # custom prometheus metrics
│   ├── metricer.go
│   ├── metrics.go
│   └── readme.md
├── proto
│   ├── blockchain.pb.go # generated code
│   └── blockchain.proto # definition
├── server
│   └── server.go
└── vendor
    ├── github.com
    ├── golang.org
    └── google.golang.org
```

## commands

```shell
# install dependencies
deep ensure

# generate code -I path to dependencies, --go_out include plugins
protoc \
  -I. \
  -I$GOPATH/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  ./proto/blockchain.proto

# build & run server
go build ./main.go
go run ./main.go

# build & run client
go build ./client/main.go
go run ./client/main.go -add=asd,888
go run ./client/main.go -list

```

## paths

```shell
export GOPATH=$HOME/go
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin:$GOROOT/bin
```