package providers_postgres

import (
	appy_driver "github.com/nfwGytautas/appy-go/driver"
)

type Provider struct {
}

func Init() (*Provider, error) {
	var err error
	provider := &Provider{}

	const connectionString = "postgres://postgres:postgres@localhost:5432/postgres"

	// Initialize the driver
	err = {{.ProjectName}}_driver.Initialize(connectionString)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

// Proxy function to the driver
func StartTransaction() (*appy_driver.Tx, error) {
	return appy_driver.StartTransaction()
}

// Proxy function to the driver
func RunTransaction(fn appy_driver.TransactionFn) error {
	return appy_driver.RunTransaction(fn)
}
