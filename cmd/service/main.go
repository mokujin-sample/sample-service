package main

import (
	"net/http"

	db "sample-service/adapter/mysql"
	"sample-service/config"
	sampleService "sample-service/implementation/service"
	sampleRepository "sample-service/repository"
	router "sample-service/transport/http"
)

func main() {
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	db, err := db.Connect(configuration.Database)
	if err != nil {
		panic(err)
	}

	repo := sampleRepository.New(db)
	svc := sampleService.NewService(repo)

	httpRouter := router.NewHTTPHandler(svc)
	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}
}
