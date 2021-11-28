dist: generate_go_code  generate_ruby_code

generate_go_code:
	protoc \
		--go_out=./pkg/proto/ --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/proto/ --go-grpc_opt=paths=source_relative \
		event_stream_filter.proto

generate_ruby_code:
	grpc_tools_ruby_protoc \
		--ruby_out=dist/ruby/lib/event-stream-filter/proto/ \
		--grpc_out=dist/ruby/lib/event-stream-filter/proto/ \
		event_stream_filter.proto
