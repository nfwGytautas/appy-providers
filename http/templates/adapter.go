package {{.DomainName}}_adapter_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.Module}}/providers"
)

func Connect(providers *providers.Providers, domain {{.DomainName}}DomainInterface) error {
	providers.Http.Api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to the {{.DomainName}} domain!")
	})

	return nil
}
