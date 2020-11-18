build:
	@protoc src/interfaces/rpcnode/rpcnode.proto --go_out=plugins=grpc:.

install:
	go get github.com/hoisie/web
	go get google.golang.org/grpc