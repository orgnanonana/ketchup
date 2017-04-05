export PATH := $(GOPATH)/bin:$(PATH)

# all proto files in subfolders
PROTO_FILES=$(shell find . -name "*.proto")

# includes for all subfolders containing proto files, e.g. `-I./path/to/proto/folder`
INCL_PROTO_DIR=$(shell find . -name "*.proto" -exec dirname {} \; | sort -u | sed -e 's/^/-I/')
INCL_WKT=-I $$GOPATH/src/github.com/golang/protobuf/ptypes/struct

PROTO_PREFIX=import_prefix_proto=github.com/octavore/ketchup/proto/

protos:
	@mkdir -p proto
	protoc $(INCL_PROTO_DIR) $(INCL_WKT) $(PROTO_FILES) $$GOPATH/src/github.com/golang/protobuf/ptypes/struct/struct.proto --go-2_out=$(PROTO_PREFIX),plugins=setter:./proto
	go run util/gots/main.go ./admin/src/js/lib/api.ts

protos_list:
	@echo $(PROTO_FILES) | tr " " "\n"

prepare:
	brew install grpc/grpc/google-protobuf
	go get -u github.com/octavore/protobuf/protoc-gen-go-2
