package {{.DomainName}}_adapter_postgres

import (
	appy_logger "github.com/nfwGytautas/appy-go/logger"
	appy_driver "github.com/nfwGytautas/appy-providers/postgres/impl"
	{{.DomainName}}_model "{{.Module}}/domains/{{.DomainName}}/model"
	{{HyphenToUnderscore .ProjectName}}_driver "github.com/nfwGytautas/homebrew-api-dash/providers/postgres/driver"
)

func (a *Adapter) Get(exampleArg string) (*{{.DomainName}}_model.Example, error) {
	tx := appy_driver.DirectConnection()
	defer tx.CommitOrRollback()

	appy_logger.Logger().Debug("Getting examples: %s", roles)

	result, err := {{HyphenToUnderscore .ProjectName}}_driver.QGetExampleById(
		tx,
		func(row appy_driver.Scannable, entry *{{.DomainName}}_model.Example) error {
			return row.Scan(
				&entry.Id,
				&entry.Name,
			)
		},
		exampleArg,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}
