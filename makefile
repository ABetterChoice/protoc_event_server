all: .build

.PHONY: build

.build: .format
	go mod tidy
	protoc --proto_path=${GOPATH}/src:${GOPATH}/src/git.code.oa.com/trpcprotocol:. --go-grpc_out=require_unimplemented_servers=false:. --trpc2grpc_out=require_unimplemented_servers=false:. --go_out=. --grpc-gateway_out=. event_server.proto
	mv -f git.code.oa.com/tencent_abtest/protocol/protoc_event_server/* ./
	rm -rf git.code.oa.com

.format:
	go mod tidy
	gofmt -w .
	goimports -w .
	golint ./...
	gonote .