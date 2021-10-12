package main

import (
	pb "adbo/timeservice/service"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	grpc "google.golang.org/grpc"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type gRPCServer struct {
	pb.UnimplementedTimeServer
}

func (g *gRPCServer) Now(ctx context.Context, in *emptypb.Empty) (*pb.Time, error) {
	log.Println("NewCourse request received")

	reply := pb.Time{
		Time: time.Now().UTC().String(),
	}

	return &reply, nil
}

func main() {
	port := 8080
	if len(os.Args) > 1 {
		parsedPort, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("Could not parse the given port!")
		}
		port = parsedPort
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Server running on port %d\n", port)

	var options []grpc.ServerOption

	grpcServer := grpc.NewServer(options...)
	pb.RegisterTimeServer(grpcServer, &gRPCServer{})
	grpcServer.Serve(lis)
}
