
protoc -I . -I third_party --go_out=src/api/v1  --go-grpc_out=src/api/v1  src/protos/v1/customer.proto
protoc -I . -I third_party --go_out=src/api/v1  --go-grpc_out=src/api/v1  src/protos/v1/api_response.proto
protoc -I . -I third_party --go_out=src/api/v1  --go-grpc_out=src/api/v1  src/protos/v1/user.proto 
protoc -I . -I third_party --grpc-gateway_out=src/api/v1 src/protos/v1/customer.proto
protoc -I . -I third_party --grpc-gateway_out=src/api/v1 src/protos/v1/user.proto
protoc -I . -I third_party --swagger_out=. src/protos/v1/customer.proto
protoc -I . -I third_party --swagger_out=. src/protos/v1/api_response.proto
protoc -I . -I third_party --swagger_out=. src/protos/v1/user.proto