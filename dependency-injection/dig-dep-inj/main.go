package main

import (
	"github.com/jpalves13/golang/dependency-injection/dig-dep-inj/config"
	"github.com/jpalves13/golang/dependency-injection/dig-dep-inj/http"
)

func main() {
	config.LoadComponents()
	person := &http.Person{
		Name: "Antonio",
		Age:  25,
	}
	config.Svc.DoAction()
	config.Pcp.DoAction(person)
}
