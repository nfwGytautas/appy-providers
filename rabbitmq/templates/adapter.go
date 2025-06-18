package {{.DomainName}}_adapter_rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"{{.Module}}/providers"
	providers_rabbitmq "{{.Module}}/providers/rabbitmq"
)

type Adapter struct {
	rabbit *providers_rabbitmq.Provider
	exampleQueue amqp091.Queue
}

func NewAdapter(providers *providers_rabbitmq.Provider) (*Adapter, error) {
	var err error
	adapter := Adapter{
		rabbit: providers.Rabbitmq,
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
