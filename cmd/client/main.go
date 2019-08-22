package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/ArangoGutierrez/pingpong/grpc/pong"
)

func main() {
	run()
}

func run() {
	conn, err := grpc.Dial("localhost:6000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewPongServiceClient(conn)
	stream, err := client.PingPongRPC(context.Background())
	waitc := make(chan struct{})

	msg := &pb.PongData{
		Msg:  "Ball",
		Ball: 0,
	}

	go func() {
		for {
			stream.Send(msg)
			time.Sleep(2 * time.Second)
			log.Println("Ping!...>", msg.Ball)
			in, err := stream.Recv()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("<...Pong!", in.Ball)
			msg.Ball = in.Ball
		}
	}()
	<-waitc
	stream.CloseSend()
}

var heya string

func SomeExportedFunc() heya{
	return "hello actions"
}
