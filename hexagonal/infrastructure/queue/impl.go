package queue

import (
	"fmt"
	"time"

	"github.com/jpalves13/golang/hexagonal/domain"
)

type ClientSqsImpl struct {
	worker domain.ClientWorker
}

func NewClientSqsImpl(worker domain.ClientWorker) ClientSqs {
	return &ClientSqsImpl{worker}
}

type ClientSqs interface {
	Do()
}

func (c *ClientSqsImpl) Do() {
	time.Sleep(3 * time.Second)
	fmt.Println("[infra] Message SQS")
	c.worker.Do()
}
