package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/sumitagrawal.0309189/simpleproject/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50053"
	defaultName = "1"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMemoryServerClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// r, err := c.DoIRemember(ctx, &pb.NumberRequest{Number: name})
	r, err := c.CheckWithMemoryServer(ctx, &pb.NumberRequest{Number: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
