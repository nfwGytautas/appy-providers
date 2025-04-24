package providers_postgres

import (
	appy_driver "github.com/nfwGytautas/appy-providers/postgres/impl"
)

type Provider struct {
}

func Init() (*Provider, error) {
	var err error
	provider := &Provider{}

	const connectionString = "postgres://postgres:postgres@localhost:5432/postgres"

	// Initialize the driver
	err = {{HyphenToUnderscore .ProjectName}}_driver.Initialize(connectionString)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
