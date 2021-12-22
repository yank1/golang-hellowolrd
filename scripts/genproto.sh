#!/usr/bin/env bash
#
# Generate all etcd protobuf bindings.
# Run from repository root directory named etcd.
#
set -e

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin:~/go/bin/

protoc -I=".:/usr/include:./internal/vender/proto" --go_out=plugins=grpc:pkg  ./api/proto/v1/helloworld.proto
protoc -I=".:/usr/include:./internal/vender/proto" --grpc-gateway_out . \
       --grpc-gateway_opt logtostderr=true \
       --grpc-gateway_opt paths=source_relative \
       --grpc-gateway_opt generate_unbound_methods=true \
      ./api/proto/v1/helloworld.proto