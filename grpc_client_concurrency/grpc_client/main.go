package main

import (
	"context"
	"flag"
	"google.golang.org/grpc/grpclog"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)

var concurrency = flag.Int("c", 1, "concurrency")

func main() {
	flag.Parse()

	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout))


	// Contact the server and print out its response.
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	//ctx := context.Background()

	var waitGroup sync.WaitGroup

	var cnt int64 = 0
	var sum int64 = 0

	for i := 0; i < *concurrency; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()// Set up a connection to the server.
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := pb.NewGreeterClient(conn)
			for {
				ctx, cancel := context.WithTimeout(context.Background(), 6 * time.Second)
				name := strings.Repeat("wait_1ms ", 1000)
				startTime := time.Now()
				_, err := c.SayHello(ctx, &pb.HelloRequest{Name: name}, grpc.FailFast(false))
				endTime := time.Now()
				dur := endTime.Sub(startTime)
				atomic.AddInt64(&cnt, 1)
				atomic.AddInt64(&sum, int64(dur))
				log.Printf("平均耗时: %v", time.Duration(sum / cnt))
				if err != nil {
					log.Printf("could not greet: %v", err)
				}
				cancel()
			}
		}()
	}

	waitGroup.Wait()
}