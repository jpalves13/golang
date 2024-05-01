package ClientAuth

import (
	"fmt"
	"time"
)

type User struct {
	Email string
	Pass  string
}

type ClientAuthImpl struct{}

func NewClientAuthImpl() ClientAuth {
	return &ClientAuthImpl{}
}

type ClientAuth interface {
	Do(user *User)
}

func (c *ClientAuthImpl) Do(user *User) {
	time.Sleep(3 * time.Second)
	fmt.Println("[App] Start auth")
	fmt.Sprintln(user.Email, " and ", user.Pass, " ClientAuthenticated")
}
