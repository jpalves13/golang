package repository

import "fmt"

type ClientRepositoryImpl struct{}

func NewClientRepositoryImpl() ClientRepository {
	return &ClientRepositoryImpl{}
}

type ClientRepository interface {
	Do()
}

func (c *ClientRepositoryImpl) Do() {
	fmt.Println("Repository done!")
}
