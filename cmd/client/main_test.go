package main_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ArangoGutierrez/pingpong/grpc/mock"
	"github.com/ArangoGutierrez/pingpong/grpc/pong"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

var msg = &pong.PongData{
	Msg:  "Ball",
	Ball: 0,
}

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stream := mock.NewMockPongService_PingPongRPCClient(ctrl)
	// set expectation on sending.
	stream.EXPECT().Send(
		gomock.Any(),
	).Return(nil)
	// Set expectation on receiving.
	stream.EXPECT().Recv().Return(msg, nil)
	stream.EXPECT().CloseSend().Return(nil)
	// Create mock for the client interface.
	rgclient := mock.NewMockPongServiceClient(ctrl)
	// Set expectation on RouteChat
	rgclient.EXPECT().PingPongRPC(
		gomock.Any(),
	).Return(stream, nil)
	if err := tespingpong(rgclient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func tespingpong(client *mock.MockPongServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.PingPongRPC(ctx)
	if err != nil {
		return err
	}
	if err := stream.Send(msg); err != nil {
		return err
	}
	if err := stream.CloseSend(); err != nil {
		return err
	}
	got, err := stream.Recv()
	if err != nil {
		return err
	}
	if !proto.Equal(got, msg) {
		return fmt.Errorf("stream.Recv() = %v, want %v", got, msg)
	}
	return nil
}
