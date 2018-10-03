#!/bin/bash

echo -n "👀 checking prerequisite: ";

if command -v protoc > /dev/null;
then
    echo "✅";
else
    echo "you need protobuf";
    exit 1;
fi

echo -n "🙏 Generate proto files: ";

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  --grpc-gateway_out=logtostderr=true:. \
  --swagger_out=logtostderr=true:. \
  api/v1/fizzbuzz/fizzbuzz.proto;

echo "✅";

echo "🏃‍ Running the fizzbuzzer";
go run fizzbuzz.go;

