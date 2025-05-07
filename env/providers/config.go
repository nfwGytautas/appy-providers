package providers_env

import (
	"os"

	"github.com/joho/godotenv"
)

type Provider struct {
}

func Init() (*Provider, error) {
	provider := &Provider{}

	// Check if .env file exists
	_, err := os.Stat(".env")
	if err != nil {
		if os.IsNotExist(err) {
			// .env file does not exist, so nothing to do
			return provider, nil
		}

		return nil, err
	}

	err = godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (p *Provider) Start() error {
	// No need to start anything for env provider
	return nil
}
