# FizzBuzzer

This project is a simple boilerplate for an API using gRPC protocol by using [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) ( also create a R-proxy to handle HTTP )

The contract of the API is handled by [Protobuf](https://developers.google.com/protocol-buffers/) interchange format 

## Context
This simple API contains only one endpoint as example which is : 
* gRPC: `GetFizzBuzz(string1, string2 string, int1, int2, limit int64)`
* HTTP: 
  * `/v1/fizzbuzz?string1=x&string2=x&int1=x&int2=x&limit=x`
  * `/v1/fizzbuzz/string1/string2/int1/int2/limit`

```text
The original fizz-buzz consists in writing all numbers from 1 to 100, and replacing :
* all multiples of 3 by “fizz”
* all multiples of 5 by “buzz”
* all multiples of 15 by “fizzbuzz”. 

The output would look like this:
“1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,fizz,...”
```

## 

## Prerequisites
I assume that you have already installed :
* Go -> [Install Go](https://golang.org/doc/install)
* Protobuf -> [Install Protobuf](https://github.com/protocolbuffers/protobuf)

## Getting started

```shell
go get -u github.com/tommy-42/go-grpc-fizzbuzzer

cd $GOPATH/src/github.com/tommy-42/go-grpc-fizzbuzzer

go get ./...

# chmod +x run.sh
./run.sh 
```