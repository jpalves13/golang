package main

import (
	"fmt"

	app "github.com/jpalves13/golang/hexagonal/application"
	"github.com/jpalves13/golang/hexagonal/domain"
	"github.com/jpalves13/golang/hexagonal/infrastructure/db"
	"github.com/jpalves13/golang/hexagonal/infrastructure/queue"
)

func main() {
	fmt.Println("Service start")
	ath := app.NewClientAuthImpl()
	db := db.NewClientDbImpl()
	wrk := domain.NewClientWorkerImpl(ath, db)
	ncs := queue.NewClientSqsImpl(wrk)
	ncs.Do()
}
