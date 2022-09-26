package main

import (
	"github.com/fokurly/avito-tech-test-task/service"
	"github.com/fokurly/avito-tech-test-task/storage/postgres"
)

func main() {
	storage, err := postgres.NewDatabase()
	if err != nil {
		panic(err)
	}

	//configKey := "storage_config"
	s := service.NewClientBalanceService(storage)
	if err := s.Run(); err != nil {
		panic(err)
	}
}
