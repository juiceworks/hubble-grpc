#!/bin/bash

SED_CMD="sed"

if [[ "$OSTYPE" == "darwin"* ]]; then
    SED_CMD="gsed"
fi

for cmd in $SED_CMD protoc curl jq; do
    command -v $cmd >/dev/null 2>&1 || { echo >&2 "Could not find `$cmd`, which is required to run this script."; exit 1; }
done

repo="farcasterxyz/hub-monorepo"
path="protobufs/schemas"
branch="main"

# Get a list of files in the schemas directory and download them.
files=$(curl -s "https://api.github.com/repos/$repo/contents/$path?ref=$branch" | jq -r '.[].download_url')
mkdir schemas
for file in $files; do
    curl -sL "$file" -o "./schemas/$(basename "$file")"
done

# Add go_package to each schema file.
for file in ./schemas/*.proto; do
    if [ -f "$file" ]; then
        $SED_CMD -i '2i option go_package = "github.com/juiceworks/hubble-grpc";' "$file"
    fi
done

# Generate the Go code.
protoc --proto_path=schemas --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative schemas/*.proto