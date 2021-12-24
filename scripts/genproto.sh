#!/usr/bin/env bash
#
# Generate all etcd protobuf bindings.
# Run from repository root directory named etcd.
#
set -ex

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin:~/go/bin/

# GRPC
protoc -I=".:./internal/vender/proto" --go_out=./pkg  ./api/proto/v1/helloworld.proto
protoc -I=".:./internal/vender/proto" --go-grpc_out=require_unimplemented_servers=false:./pkg  ./api/proto/v1/helloworld.proto

# GRPC GW

protoc -I=".:./internal/vender/proto" --grpc-gateway_out . \
       --grpc-gateway_opt logtostderr=true \
       --grpc-gateway_opt paths=source_relative \
       --grpc-gateway_opt generate_unbound_methods=true \
      ./api/proto/v1/helloworld.proto
mv ./api/proto/v1/helloworld.pb.gw.go ./pkg/helloworld

# OpenAPI

protoc -I=".:./internal/vender/proto" --openapiv2_out=. ./api/proto/v1/helloworld.proto
mv ./api/proto/v1/helloworld.swagger.json ./api/swagger/v1/helloworld.swagger.json
cp ./api/swagger/v1/helloworld.swagger.json assets/swagger-ui/helloworld.swagger.json 

# Typescript

cd ./api/proto/v1/
protoc -I=".:./../../../internal/vender/proto" --grpc-gateway-ts_out=. helloworld.proto
cd ../../../
