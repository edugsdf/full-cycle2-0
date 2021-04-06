package main

import (
	"log"
	"net"

	"github.com/edugsdf/full-cycle2-0/tree/main/docker/golang/go-aluno/pb/pb"
	"github.com/edugsdf/full-cycle2-0/tree/main/docker/golang/go-aluno/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not conect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not server: %v", err)
	}
}
