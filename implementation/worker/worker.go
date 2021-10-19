package objectworker

import (
	svc "sample-service"

	"github.com/streadway/amqp"
)

type worker struct {
	objectRepository   svc.Repository
	currentObject      *svc.Object
	produceQueueNotify *amqp.Channel
	objectsNotify      chan svc.Object
}

func NewWorker(objectRepo svc.Repository, produceQueueNotify *amqp.Channel) svc.ObjectWorker {
	return &worker{
		objectRepository:   objectRepo,
		produceQueueNotify: produceQueueNotify,
		objectsNotify:      make(chan svc.Object),
	}
}

func (w *worker) ProcessObjects() {

}

func (w *worker) ProcessNotifications() {

}
