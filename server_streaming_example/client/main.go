package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "GolandProjects/grpc-tutorial/server_streaming_example/proto"
)

func main() {
	// dial to server
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Println("Error connecting to gRPC server: ", err.Error())
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Falied to connect: %v\n", err)
		}
	}(conn)

	// create the stream
	client := pb.NewStreamingServiceClient(conn)

	req := pb.DataRequest{Id: "123"}
	stream, err := client.GetDataStreaming(context.Background(), &req)
	if err != nil {
		panic(err) // dont use panic in your real project
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return
		} else if err == nil {
			valStr := fmt.Sprintf("Response\n Part: %d \n Val: %s", resp.Part, resp.Buffer)
			log.Println(valStr)
		}

		if err != nil {
			panic(err) // dont use panic in your real project
		}

	}
}
