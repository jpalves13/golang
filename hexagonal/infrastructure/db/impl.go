package db

import (
	"fmt"
	"time"
)

type ClientDbImpl struct{}

func NewClientDbImpl() ClientDb {
	return &ClientDbImpl{}
}

type ClientDb interface {
	Do()
}

func (c *ClientDbImpl) Do() {
	time.Sleep(3 * time.Second)
	fmt.Println("[infra] Save DB!")
}
