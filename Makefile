build:
	@protoc src/interfaces/rpcnode/rpcnode.proto --go_out=plugins=grpc:.

install:
	go install github.com/hoisie/web
	go install google.golang.org/grpc
	go install github.com/hoisie/web
	go install google.golang.org/grpc/encoding/proto