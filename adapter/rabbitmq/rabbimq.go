package rabbitmq

import "github.com/streadway/amqp"

func Connect(dsn string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func ConsumeQueue(connection *amqp.Connection, queueName string) (<-chan amqp.Delivery, error) {
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}
	err = channel.Qos(1, 0, false)
	if err != nil {
		return nil, err
	}

	delivery, err := channel.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return delivery, nil
}

func ProduceQueue(connection *amqp.Connection, queueName string) (channel *amqp.Channel, err error) {
	channel, err = connection.Channel()
	if err != nil {
		return nil, err
	}
	args := make(amqp.Table)
	args["x-max-priority"] = int64(10)
	_, err = channel.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		return nil, err
	}

	return channel, nil
}
