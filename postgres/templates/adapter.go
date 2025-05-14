package {{.DomainName}}_adapter_postgres

import (
	"{{.Module}}/providers"
)

type Adapter struct {
}

func NewAdapter(providers *providers.Providers) *Adapter {
	return &Adapter{}
}
