# Matomat-server

[![Maintainability](https://api.codeclimate.com/v1/badges/7513a48ef0f203220eae/maintainability)](https://codeclimate.com/github/omgwtflaserguns/matomat-server/maintainability)

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

