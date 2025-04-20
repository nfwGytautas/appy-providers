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
		return nil, err
	}

	err = godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return provider, nil
}
