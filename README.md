
# Matomat-server

go grpc und grpc-web server

## setup project

1. install protobuf and protoc-go
2. git clone --recursive
3. dep ensure
4. go generate

## Add dependency

1. Add import in go code
2. dep ensure

## Update contracts

1. git submodule update --remote --merge
2. go generate

