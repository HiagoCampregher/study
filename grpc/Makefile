.PHONY: generate test server client

generate:
	protoc \
		--go_out=. --go_opt=module=study/grpc \
		--go-grpc_out=. --go-grpc_opt=module=study/grpc \
		proto/course_category.proto

test:
	go test ./...

server:
	go run ./cmd/grpcServer

client:
	go run ./cmd/grpcClient
