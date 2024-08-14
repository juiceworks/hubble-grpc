# Hubble gRPC Bindings for Golang

Basic gRPC bindings for interacting with [Farcaster Hubs](https://github.com/farcasterxyz/hub-monorepo).

To install in a Go project:

```bash
go get github.com/juiceworks/hubble-grpc
```

To generate updated bindings:

```bash
git clone https://github.com/juiceworks/hubble-grpc
cd hubble-grpc
chmod +x main.sh
./main.sh
```

## Usage

You can use these bindings with public hubs – see [this list](https://foss.farchiver.xyz/hubs/list). See [example/example.go](example/example.go) for basic usage example which prints out new casts.

## About

The `.proto` schemas are sourced from the [Hub Monorepo](https://github.com/farcasterxyz/hub-monorepo/tree/main/protobufs/schemas). `main.sh` adds a `go_package` option to each schema file beneath its syntax declaration:

```proto
option go_package = "github.com/juiceworks/hubble-grpc";
```

Bindings are then generated with `protoc`:

```bash
protoc --proto_path=schemas --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative schemas/*.proto
```
