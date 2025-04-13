package {{.DomainName}}_adapter_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ports "{{.PackageRoot}}/domains/{{.DomainName}}/ports"
	providers_http "{{.PackageRoot}}/providers/http"
)

type {{.DomainName}}HttpAdapter struct {
	Usecase ports.Get{{.DomainName}}InputPort
}

func (a *{{.DomainName}}HttpAdapter) Setup(server *providers_http.HttpProvider) {
	server.Root.GET("/{{.DomainName}}/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		tenant, err := a.Usecase.Execute(ports.Get{{.DomainName}}Command{ID: id})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, tenant)
	})
}
