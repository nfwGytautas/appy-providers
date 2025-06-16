package {{.DomainName}}_adapter_rabbitmq

import (
	"net/http"

	"github.com/rabbitmq/amqp091-go"
	"{{.Module}}/providers"
)

func (a *Adapter) StartReceiving() error {
	msgs, err := ch.Consume(
		a.exampleQueue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	// Start a goroutine to process messages
	go func() {
		for msg := range msgs {
			// Process the message
		}
	}()

	return nil
}
