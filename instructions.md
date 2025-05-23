1. install go dependencies
   ```
    go mod tidy
   ```
2. generate proto
   ```
    protoc \   
    --go_out=./microservices/pb --go_opt=paths=source_relative \
    --go-grpc_out=./microservices/pb --go-grpc_opt=paths=source_relative \
    api/serviceA/serviceA.proto \
    api/serviceB/serviceB.proto
   ```