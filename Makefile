generate_protos:
	protoc \
		--go_out=./pkg/proto/ --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/proto/ --go-grpc_opt=paths=source_relative \
		event-stream-filter.proto



