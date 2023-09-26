run-gateway:
	go run main.go

build-gateway:
	protoc -I http/user/proto --go_out=gateway/http/user --go-grpc_out=http/user http/user/proto/*.proto
	protoc -I http/petition/proto --go_out=gateway/http/petition --go-grpc_out=http/petition http/petition/proto/*.proto


clear-gateway:
	rm http/petition/pb/*.go
	rm http/user/pb/*.go