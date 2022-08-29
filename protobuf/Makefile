OS_TYPE := $(shell echo `uname`)
OS_ARCH := $(shell echo `uname -m`)

ifeq ($(OS_TYPE), Darwin)
	OS_TYPE := osx
else ifeq ($(OS_TYPE), Linux)
	OS_TYPE := linux
endif

ifeq ($(OS_ARCH), arm64)
	OS_ARCH := aarch_64
endif

PROTOC_VERSION := 21.5
PROTOC_ARCH := $(OS_TYPE)-$(OS_ARCH)
PROTOC_ZIP := protoc-$(PROTOC_VERSION)-$(PROTOC_ARCH).zip
PROTOC := ./protoc/v$(PROTOC_VERSION)/bin/protoc
PROTOC_OUT := ../go/protobuf/protoc
PROTOC_GO_OPT := paths=source_relative
PROTOC_GO_GRPC_OPT := paths=source_relative,require_unimplemented_servers=false

.PHONY: buf
buf:
	buf generate blogapis

.PHONY: protoc
protoc:
	$(PROTOC) \
		-I protoc/v$(PROTOC_VERSION)/include \
		-I blogapis \
		--go_out=$(PROTOC_OUT) \
		--go_opt=$(PROTOC_GO_OPT) \
		--go-grpc_out=$(PROTOC_OUT) \
		--go-grpc_opt=$(PROTOC_GO_GRPC_OPT) \
		blog/v1/post.proto blog/v1/user.proto

.PHONY: clean-protoc
clean-protoc:
	rm -rf protoc/v$(PROTOC_VERSION)

.PHONY: install-protoc
install-protoc: clean-protoc
	curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_ZIP)
	mkdir -p ./protoc/v$(PROTOC_VERSION)
	unzip $(PROTOC_ZIP) -d ./protoc/v$(PROTOC_VERSION)
	rm $(PROTOC_ZIP)
