#!/bin/bash

# install required packages:
#
# go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
# go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
# go get -u github.com/golang/protobuf/protoc-gen-go
# go get -u github.com/grpc-ecosystem/grpc-gateway/... 

# grpc server
protoc -I. -I ../../../../../src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,import_path=contract:. ./*.proto

# http gateway
protoc -I. -I ../../../../../src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,request_context=true,import_path=contract:. ./*.proto

# swagger documentation
protoc -I. -I ../../../../../src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:./../ ./clients.proto
protoc -I. -I ../../../../../src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:./../ ./authz.proto
