run-gateway:
	go run gateway/main.go

build-gateway:
	protoc -I gateway/http/user/proto --go_out=gateway/http/user --go-grpc_out=gateway/http/user gateway/http/user/proto/*.proto
	protoc -I gateway/http/petition/proto --go_out=gateway/http/petition --go-grpc_out=gateway/http/petition gateway/http/petition/proto/*.proto


clear-gateway:
	rm gateway/http/petition/pb/*.go
	rm gateway/http/user/pb/*.go