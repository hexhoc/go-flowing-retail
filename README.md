# go-flowing-retail

<!-- TODO: https://github.com/morzhanov/go-realworld -->
<!-- TODO: go-coffeeshop -->
## Compile proto files
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. .\internal\pb\*.proto
