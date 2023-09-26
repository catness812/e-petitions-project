run-gateway:
	go run main.go

build-gateway:
	protoc -I ./internal/user/proto --go_out=./internal/user --go-grpc_out=./internal/user ./internal/user/proto/*.proto
	protoc -I ./internal/petition/proto --go_out=./internal/petition --go-grpc_out=./internal/petition ./internal/petition/proto/*.proto
	protoc -I ./internal/security/proto --go_out=./internal/security --go-grpc_out=./internal/security ./internal/security/proto/*.proto


clear-gateway:
	rm internal/petition/pb/*.go
	rm internal/user/pb/*.go