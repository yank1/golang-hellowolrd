#!/usr/bin/env bash
#
# Generate all etcd protobuf bindings.
# Run from repository root directory named etcd.
#
set -e

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin:~/go/bin/

protoc -I . -I /usr/include --go_out=plugins=grpc:pkg ./api/proto/*/*.proto