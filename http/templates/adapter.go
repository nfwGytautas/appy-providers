package {{.DomainName}}_adapter_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.Module}}/providers"
)

type Domain interface {
	// Describe the methods that you want to provide HTTP access to
}

func Connect(providers *providers.Providers, domain Domain) error {
	providers.Http.Api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to the {{.DomainName}} domain!")
	})

	return nil
}
