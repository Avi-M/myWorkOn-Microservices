package main

import (
	"demo_grpc/domain"
	"demo_grpc/impl"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	netListener := getNetListener(7000)
	grpcServer := grpc.NewServer()

	repositoryServiceImpl := impl.NewRepositoryServiceGrpcImpl()
	domain.RegisterRepositoryServiceServer(grpcServer, repositoryServiceImpl)

	// start the server
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
