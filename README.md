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

## Resources

1. To find public hubs you can use `hubble-grpc` with, see [this list of public hubs](https://foss.farchiver.xyz/hubs/list).
2. To learn about the Farcaster Hub gRPC API, see the [Farcaster Hub API Docs](https://docs.farcaster.xyz/reference/hubble/architecture).
3. For a basic example which prints out new casts, see [example/example.go](example/example.go).

## About

The `.proto` schemas are sourced from the [Hub Monorepo](https://github.com/farcasterxyz/hub-monorepo/tree/main/protobufs/schemas). `main.sh` adds a `go_package` option to each schema file beneath its syntax declaration:

```proto
option go_package = "github.com/juiceworks/hubble-grpc";
```

Bindings are then generated with `protoc`:

```bash
protoc --proto_path=schemas --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative schemas/*.proto
```
