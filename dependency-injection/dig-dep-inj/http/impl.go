package http

import (
	"fmt"

	"github.com/jpalves13/golang/dependency-injection/dig-dep-inj/auth"
)

type Person struct {
	Name string
	Age  int
}

type ClientHttpImpl struct {
	clientAuth auth.ClientAuth
}

func NewClientHttpImpl(clientAuth auth.ClientAuth) ClientHttp {
	return &ClientHttpImpl{clientAuth}
}

type ClientHttp interface {
	DoAction(request *Person)
}

func (c *ClientHttpImpl) DoAction(request *Person) {
	c.clientAuth.Do()
	fmt.Println("Http done!")
}
