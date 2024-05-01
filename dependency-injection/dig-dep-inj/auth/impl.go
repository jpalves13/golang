package auth

import "fmt"

type ClientAuthImpl struct{}

func NewClientAuthImpl() ClientAuth {
	return &ClientAuthImpl{}
}

type ClientAuth interface {
	Do()
}

func (c *ClientAuthImpl) Do() {
	fmt.Println("auth done!")
}
