package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/ArangoGutierrez/PingPong/grpc/pong"
)

func main() {
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
