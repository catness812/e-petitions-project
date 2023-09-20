run-gateway:
	go run gateway/main.go

build-gateway:
	#protoc -I /gateway/http/petition/proto --go_out=./petition --go-grpc_out=.gateway/http/petition petition/proto/*.proto
	protoc -I /gateway/http/user/proto --go_out=./user --go-grpc_out=.gateway/http/user user/proto/*.proto

clear-gateway:
	rm gateway/http/petition/pb
	rm gateway/http/user/pb