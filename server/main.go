package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sumitagrawal.0309189/simpleproject/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50053"
)

type server struct {
	pb.UnimplementedMemoryServerServer
}

func itemExists(item string, arr []string) bool {
	var memoryCapacity = 3
	for i := len(arr); i > 0 && memoryCapacity > 0; i-- {
		if arr[i-1] == item {
			return true
		}
		memoryCapacity--
	}
	return false
}

var memory = make([]string, 0)

func (s *server) CheckWithMemoryServer(ctx context.Context, in *pb.NumberRequest) (*pb.DoIRemember, error) {

	for i := 0; i < len(memory); i++ {
		log.Printf("server_memory %v", memory[i])
	}

	incomingNumber := in.GetNumber()

	log.Printf("Received: %v", incomingNumber)

	if itemExists(incomingNumber, memory) {
		return &pb.DoIRemember{Message: "Yes I remember " + incomingNumber}, nil
	}
	memory = append(memory, incomingNumber)
	return &pb.DoIRemember{Message: "It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number, I will try to remember " + incomingNumber}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMemoryServerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
