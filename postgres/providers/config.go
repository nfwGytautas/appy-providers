package providers_postgres

import (
	{{HyphenToUnderscore .Config.ProjectName}}_driver "{{.Config.Module}}/providers/postgres/driver"
)

type Provider struct {
}

func Init() (*Provider, error) {
	var err error
	provider := &Provider{}

	const connectionString = "postgres://postgres:postgres@localhost:5432/postgres"

	// Initialize the driver
	err = {{HyphenToUnderscore .Config.ProjectName}}_driver.Initialize(connectionString)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (p *Provider) Start() error {
	// postgres driver does not require any specific start logic
	return nil
}
