package main

import (
	"context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"os"
)

const (
	port = "127.0.0.1:50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//p, ok := peer.FromContext(ctx)
	//if ok {
	//	fmt.Printf("p=%+v\n", p)
	//} else {
	//	fmt.Println("not ok!")
	//}
	//log.Printf("Received: %v", in.Name)
	//if strings.HasPrefix(in.GetName(), "wait_1ms") {
	//	time.Sleep(1 * time.Millisecond)
	//} else {
	//	time.Sleep(100 * time.Millisecond)
	//}

	message := "Hello "
	for i := 0; i < 1000; i++ {
		message += "Hello"
	}
	return &pb.HelloReply{Message: message}, nil
}

func main() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}