package config

import (
	"github.com/jpalves13/golang/dependency-injection/dig-dep-inj/auth"
	"github.com/jpalves13/golang/dependency-injection/dig-dep-inj/http"
	"github.com/jpalves13/golang/dependency-injection/dig-dep-inj/repository"
	"github.com/jpalves13/golang/dependency-injection/dig-dep-inj/service"
	"go.uber.org/dig"
)

var digContainer *dig.Container
var Svc service.ClientService
var Pcp http.ClientHttp

func LoadComponents() {
	provideComponents()
	initComponents()
}

func initComponents() {
	digContainer.Invoke(
		func(s service.ClientService) {
			Svc = s
		},
	)
	digContainer.Invoke(
		func(s http.ClientHttp) {
			Pcp = s
		},
	)
}

func provideComponents() {
	digContainer = dig.New()

	digContainer.Provide(
		func() repository.ClientRepository {
			return repository.NewClientRepositoryImpl()
		},
	)
	digContainer.Provide(service.NewClientServiceImpl)

	digContainer.Provide(
		func() auth.ClientAuth {
			return auth.NewClientAuthImpl()
		},
	)
	digContainer.Provide(http.NewClientHttpImpl)
}
