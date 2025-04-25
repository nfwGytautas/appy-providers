package providers_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appy_http "github.com/nfwGytautas/appy-go/http"
)

type Provider struct {
	Server *appy_http.Server

	Root    *gin.RouterGroup
	Api     *gin.RouterGroup
	App     *gin.RouterGroup
	Public  *gin.RouterGroup
	Private *gin.RouterGroup
}

func Init() (*Provider, error) {
	var err error
	provider := &Provider{}

	// Initialize HTTP server
	config := appy_http.HttpConfig{
		ErrorMapper: &ErrorMapper{},
		Address:     "127.0.0.1:8080",
	}

	provider.Server, err = appy_http.InitializeHTTP(&config)
	if err != nil {
		return nil, err
	}

	// Initialize root groups
	provider.Root = provider.Server.Root()
	provider.App = provider.Root.Group("/app")
	provider.Api = provider.Root.Group("/api")
	provider.Public = provider.Api.Group("/")
	provider.Private = provider.Api.Group("/")

	// Setup middleware
	// ...

	// It is suggested to leave this endpoint for automated tools
	provider.Public.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusNoContent)
	})

	return provider, nil
}
