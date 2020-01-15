package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	GetArea "./GetArea/proto/GetArea"
)

type GetArea struct{}

func (e *GetArea) Handle(ctx context.Context, msg *GetArea.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *GetArea.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
