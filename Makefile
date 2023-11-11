proto_security:
	protoc -I security_service/internal/proto --go_out=security_service/internal/ --go-grpc_out=require_unimplemented_servers=false:security_service/internal security_service/internal/proto/*.proto
proto_security_gateway:
	protoc -I gateway/internal/security/proto --go_out=gateway/internal/security --go-grpc_out=gateway/internal/security gateway/internal/security/proto/*.proto

proto_user:
	protoc -I user_service/internal/proto --go_out=user_service/internal/ --go-grpc_out=require_unimplemented_servers=false:user_service/internal user_service/internal/proto/*.proto
proto_user_gateway:
	protoc -I gateway/internal/user/proto --go_out=gateway/internal/user --go-grpc_out=gateway/internal/user gateway/internal/user/proto/*.proto

mail-docker-build:
	docker build -t e-petitions-mail:1.0 ./mail_serviceproto_security:
	protoc -I security_service/internal/proto --go_out=security_service/internal/ --go-grpc_out=require_unimplemented_servers=false:security_service/internal security_service/internal/proto/*.proto

proto_file:
	protoc -I file_service/internal/proto --go_out=file_service/internal/ --go-grpc_out=require_unimplemented_servers=false:file_service/internal file_service/internal/proto/*.proto

build_container_images:
	docker build -t e-petitions-gateway:1.0 ./gateway && \
	docker build -t e-petitions-mail:1.0 ./mail_service && \
	docker build -t e-petitions-petition:1.0 ./petition_service && \
	docker build -t e-petitions-security:1.0 ./security_service  && \
	docker build -t e-petitions-user:1.0 ./user_service