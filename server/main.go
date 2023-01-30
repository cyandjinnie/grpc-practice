package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/cyandjinnie/grpc-practice/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.String("port", "50051", "port on which the server is hosted")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("- request %v", in.GetName())
	return &pb.HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("- listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
