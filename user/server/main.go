package main

import (
	pb "GolandProjects/grpc-tutorial/user/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.UserServiceServer
}

func main() {
	list, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalln("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &Server{})

	log.Println("start server")

	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
