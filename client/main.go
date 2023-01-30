package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	pb "github.com/cyandjinnie/grpc-practice/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultname = "world"
)

var (
	addr = flag.String("addr", "localhost", "the address to connect to")
	port = flag.String("port", "50051", "port on which the server is hosted")
	name = flag.String("name", defaultname, "name to greet")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr+":"+*port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Reply: %s", r.GetMessage())
}
