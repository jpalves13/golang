package main

import (
	"github.com/jpalves13/golang/dependency-injection/no-dep-inj/repository"
	"github.com/jpalves13/golang/dependency-injection/no-dep-inj/service"
)

func main() {
	// create repository
	repo := repository.NewClientRepositoryImpl()

	// create service with repository as parameter
	svc := service.NewClientServiceImpl(repo)

	// perform actions
	svc.DoAction()
}
