package service

import (
	"fmt"

	"github.com/jpalves13/golang/dependency-injection/no-dep-inj/repository"
)

type ClientServiceImpl struct {
	clientRepository repository.ClientRepository
}

func NewClientServiceImpl(clientRepository repository.ClientRepository) ClientService {
	return &ClientServiceImpl{clientRepository}
}

type ClientService interface {
	DoAction()
}

func (c *ClientServiceImpl) DoAction() {
	c.clientRepository.Do()
	fmt.Println("Service done!")
}
