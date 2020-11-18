package rpcnode

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Init(port int) {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	RegisterRpcNodeServer(grpcServer, &server{})
	grpcServer.Serve(listener)

}

type server struct {
}

func (s *server) Sync(ctx context.Context, req *RpcNodeRequest) (*RpcNodeResponse, error) {
	return &RpcNodeResponse{Result: "test"}, nil
}
