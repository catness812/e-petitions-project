proto_security:
	protoc -I security_service/internal/proto --go_out=security_service/internal/ --go-grpc_out=require_unimplemented_servers=false:security_service/internal security_service/internal/proto/*.proto
proto_security_gateway:
	protoc -I gateway/internal/security/proto --go_out=gateway/internal/security --go-grpc_out=gateway/internal/security gateway/internal/security/proto/*.proto
	

mail-docker-build:
	docker build -t e-petitions-mail:1.0 ./mail_service
