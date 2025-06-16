package {{.DomainName}}_adapter_rabbitmq

import (
	"context"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func (a *Adapter) SendMessage(body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := a.rabbit.Ch.PublishWithContext(ctx,
		"",                  // exchange
		a.exampleQueue.Name, // routing key
		false,               // mandatory
		false,               // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
