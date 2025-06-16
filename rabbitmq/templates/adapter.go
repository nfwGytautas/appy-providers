package {{.DomainName}}_adapter_rabbitmq

import (
	"net/http"

	"github.com/rabbitmq/amqp091-go"
	"{{.Module}}/providers"
)

type Adapter struct {
	rabbit *providers.RabbitMQ

	exampleQueue *amqp091.Queue
}

func NewAdapter(providers *providers.Providers) (*Adapter, error) {
	var err error
	adapter := Adapter{
		rabbit: providers.RabbitMQ,
	}

	// Declare a queue for example usage
	adapter.exampleQueue, err = adapter.rabbit.Ch.QueueDeclare(
		"example_queue",
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		return nil, err
	}

	return &adapter, nil
}
