package {{.DomainName}}_adapter_rabbitmq

import (
	"net/http"

	"github.com/rabbitmq/amqp091-go"
	"{{.Module}}/providers"
)

func (a *Adapter) SendMessage(body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = a.Ch.PublishWithContext(ctx,
		"",                     // exchange
		a.exampleQueue.Name,   // routing key
		false,                  // mandatory
		false,                  // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err.Error()
	}

	return nil
}
