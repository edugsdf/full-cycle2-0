package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/edugsdf/full-cycle2-0/tree/main/docker/golang/go-aluno/pb/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v ", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	//AddUser(client) //<--- Unário
	//AddUserVerbose(client)  //<--- Server Stream
	//AddUsers(client) //<--- Cliente Stream
	AddUserStreamBoth(client) //<--- Stream bidirecional
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "88",
		Name:  "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v ", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "88",
		Name:  "Joao",
		Email: "j@j.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v ", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receiva the menssage: %v ", err)
		}

		fmt.Println("Status = ", stream.Status, " - ", stream.User)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "ED-01",
			Name:  "Eduardo 01 ",
			Email: "edu@edu.com",
		},
		&pb.User{
			Id:    "ED-02",
			Name:  "Eduardo 02 ",
			Email: "edu02@edu.com",
		},
		&pb.User{
			Id:    "ED-03",
			Name:  "Eduardo 03",
			Email: "edu03@edu.com",
		},
		&pb.User{
			Id:    "ED-04",
			Name:  "Eduardo 04",
			Email: "edu04@edu.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Erro creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Erro receiving response: %v", err)
	}
	fmt.Println("Resposta: ", res)

}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Erro creating request: %v", err)
	}

	//--------------
	reqs := []*pb.User{
		&pb.User{
			Id:    "ED-01",
			Name:  "Eduardo 01 ",
			Email: "edu@edu.com",
		},
		&pb.User{
			Id:    "ED-02",
			Name:  "Eduardo 02 ",
			Email: "edu02@edu.com",
		},
		&pb.User{
			Id:    "ED-03",
			Name:  "Eduardo 03",
			Email: "edu03@edu.com",
		},
		&pb.User{
			Id:    "ED-04",
			Name:  "Eduardo 04",
			Email: "edu04@edu.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.GetName())
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}
			fmt.Printf("Recebendo user: %v, com status: %v \n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
