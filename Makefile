.PHONY: proto

all: tidy proto test

tidy:
	go mod tidy

proto:
	protoc -I pkg/proto/ pkg/proto/drlm.proto --go_out=plugins=grpc:pkg/proto

test:
	go test -cover ./...
