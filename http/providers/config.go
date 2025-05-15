package providers_http

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TODO: You could set this from an environment variable if needed
const address = "127.0.0.1:3000"

type Provider struct {
	Server *gin.Engine

	Root    *gin.RouterGroup
	Api     *gin.RouterGroup
	Public  *gin.RouterGroup
	Private *gin.RouterGroup
}

func Init() (*Provider, error) {
	provider := &Provider{}

	// Initialize HTTP server
	provider.Server = gin.Default()
	provider.Server.Use(cors.Default())

	// Initialize root groups
	provider.Root = provider.Server.Group("/")
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

func (p *Provider) Start() error {
	// Start the HTTP server
	go func() {
		p.Server.Run(address)
	}()

	return nil
}
