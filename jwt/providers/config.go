package providers_jwt

import (
	"time"

	appy_jwt "github.com/nfwGytautas/appy-providers/jwt/impl"
)

// TODO: Create actual key and API keys, potential to move to env variables
const secret = "xxxx-xxxx-xxxx-xxxx-xxxx"

type Claims = appy_jwt.Claims

type Provider struct {
	appy_jwt.JwtConfig
}

func Init() (*Provider, error) {
	provider := &Provider{}

	provider.JwtConfig = appy_jwt.NewJwtConfig(secret, time.Minute*5, time.Hour*24*7)

	return provider, nil
}

func (p *Provider) Start() error {
	// jwt driver does not require any specific start logic
	return nil
}
