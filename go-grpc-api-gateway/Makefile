proto:
	protoc --proto_path=pkg/auth/pb --go_out=pkg/auth/pb --go_opt=paths=source_relative \
	--go-grpc_out=pkg/auth/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pkg/auth/pb --grpc-gateway_opt=paths=source_relative \
	pkg/auth/pb/*.proto

	protoc --proto_path=pkg/order/pb --go_out=pkg/order/pb --go_opt=paths=source_relative \
	--go-grpc_out=pkg/order/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pkg/order/pb --grpc-gateway_opt=paths=source_relative \
	pkg/order/pb/*.proto

	protoc --proto_path=pkg/product/pb --go_out=pkg/product/pb --go_opt=paths=source_relative \
	--go-grpc_out=pkg/product/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pkg/product/pb --grpc-gateway_opt=paths=source_relative \
	pkg/product/pb/*.proto

server:
	go run cmd/main.go