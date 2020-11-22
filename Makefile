build:
	@protoc src/interfaces/rpcnode/rpcnode.proto --go_out=plugins=grpc:.

	go test || echo "error"

install:
	go get github.com/hoisie/web
	go get google.golang.org/grpc
	go get github.com/hoisie/web
	go get google.golang.org/grpc/encoding/proto
	go get github.com/sirupsen/logrus
	go get github.com/patrickmn/go-cache