package providers_rabbitmq

import "github.com/rabbitmq/amqp091-go"

type Provider struct {
	conn *amqp091.Connection
	Ch   *amqp091.Channel
}

func Init() (*Provider, error) {
	var err error
	provider := &Provider{}

	// TODO: Enter your RabbitMQ connection details here
	provider.conn, err = amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	provider.Ch, err = provider.conn.Channel()
	if err != nil {
		provider.conn.Close()
		return nil, err
	}

	return provider, nil
}

func (p *Provider) Start() error {
	return nil
}
