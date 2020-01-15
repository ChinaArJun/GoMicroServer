package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	GetArea "./GetArea/proto/GetArea"
)

type GetArea struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetArea) Call(ctx context.Context, req *GetArea.Request, rsp *GetArea.Response) error {
	log.Log("Received GetArea.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetArea) Stream(ctx context.Context, req *GetArea.StreamingRequest, stream GetArea.GetArea_StreamStream) error {
	log.Logf("Received GetArea.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&GetArea.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetArea) PingPong(ctx context.Context, stream GetArea.GetArea_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&GetArea.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
