# go-flowing-retail

## Compile proto files
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. .\internal\pb\*.proto
