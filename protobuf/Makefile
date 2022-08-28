PROTOC := ./bin/protoc
PROTOC_OUT := ../go/protobuf/protoc
PROTOC_GO_OPT := paths=source_relative
PROTOC_GO_GRPC_OPT := paths=source_relative,require_unimplemented_servers=false

.PHONY: buf
buf:
	buf generate blogapis

.PHONY: protoc
protoc:
	$(PROTOC) \
		-I include \
		-I blogapis \
		--go_out=$(PROTOC_OUT) \
		--go_opt=$(PROTOC_GO_OPT) \
		--go-grpc_out=$(PROTOC_OUT) \
		--go-grpc_opt=$(PROTOC_GO_GRPC_OPT) \
		blog/v1/post.proto blog/v1/user.proto
