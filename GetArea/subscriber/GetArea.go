package subscriber

import (
	getArea "GoMicroServer/GetArea/proto/GetArea"
	"context"
	"github.com/micro/go-micro/util/log"
)

type GetArea struct{}

func (e *GetArea) Handle(ctx context.Context, msg *getArea.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *getArea.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
