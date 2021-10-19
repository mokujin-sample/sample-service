package main

import (
	db "sample-service/adapter/mysql"
	"sample-service/adapter/rabbitmq"
	"sample-service/config"
	objectworker "sample-service/implementation/worker"
	objectRepository "sample-service/repository"

	"github.com/robfig/cron/v3"
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

	rabbitConnection, err := rabbitmq.Connect(configuration.RabbitMQ.DSN)
	if err != nil {
		panic(err)
	}
	defer rabbitConnection.Close()

	produceQueueNotify, err := rabbitmq.ProduceQueue(rabbitConnection, configuration.RabbitMQ.NotifyQueue)
	if err != nil {
		panic(err)
	}

	objectRepo := objectRepository.New(db)

	done := make(chan bool, 1)

	worker := objectworker.NewWorker(objectRepo, produceQueueNotify)

	go worker.ProcessNotifications()

	c := cron.New(cron.WithChain(cron.SkipIfStillRunning(cron.DiscardLogger)))
	c.Start()
	defer c.Stop()
	c.AddFunc("@every 30s", func() {
		worker.ProcessObjects()
	})

	<-done
}
