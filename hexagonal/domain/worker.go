package domain

import (
	"fmt"

	app "github.com/jpalves13/golang/hexagonal/application"
	db "github.com/jpalves13/golang/hexagonal/infrastructure/db"
)

type ClientWorkerImpl struct {
	auth app.ClientAuth
	db   db.ClientDb
}

func NewClientWorkerImpl(auth app.ClientAuth, db db.ClientDb) ClientWorker {
	return &ClientWorkerImpl{auth, db}
}

type ClientWorker interface {
	Do()
}

func (c *ClientWorkerImpl) Do() {
	fmt.Println("[Domain] Process the message")
	user := c.mock()
	c.auth.Do(user)
	c.db.Do()
}

func (c *ClientWorkerImpl) mock() *app.User {
	return &app.User{
		Email: "mock@mock.com",
		Pass:  "123456",
	}
}
