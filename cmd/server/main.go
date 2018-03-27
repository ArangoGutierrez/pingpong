package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/ArangoGutierrez/PingPong/grpc/pong"
)

type PongServer struct {
}

func (ps *PongServer) PingPongRPC(stream pb.PongService_PingPongRPCServer) error {
	log.Println("Started stream")
	for {
		in, err := stream.Recv()
		log.Println("Ping ...>  %v", in.Ball)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println("<...Pong", in.Ball)
		in.Ball++
		err = stream.Send(in)
		if err != nil {
			return err
		}
	}
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterPongServiceServer(grpcServer, &PongServer{})

	l, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Listening on tcp://localhost:6000")
	grpcServer.Serve(l)
}
