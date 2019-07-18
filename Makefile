.PHONY: proto

all: proto test
proto:
	protoc -I pkg/proto/ pkg/proto/drlm.proto --go_out=plugins=grpc:pkg/proto

test:
	go test -cover ./...
