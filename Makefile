.PHONY: proto

all: proto
proto:
	protoc -I proto/ proto/drlm.proto --go_out=plugins=grpc:proto
