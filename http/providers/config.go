package providers_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appy_http "github.com/nfwGytautas/appy-go/http"
)

// TODO: Create actual key and API keys, potential to move to env variables
const APISecret = "xxxx-xxxx-xxxx-xxxx-xxxx"
const APIKey = "xxxx-xxxx-xxxx-xxxx-xxxx"

type HttpProvider struct {
	Server *appy_http.Server

	Root    *gin.RouterGroup
	Api     *gin.RouterGroup
	App     *gin.RouterGroup
	Public  *gin.RouterGroup
	Private *gin.RouterGroup

	jwt    appy_http.JwtAuth
	appKey appy_http.ApiKeyMiddlewareProvider
}

func Init() (*HttpProvider, error) {
	var err error
	provider := &HttpProvider{}

	// Initialize HTTP server
	config := appy_http.HttpConfig{
		ErrorMapper: &ErrorMapper{},
		Address:     "127.0.0.1:8080",
	}

	provider.Server, err = appy_http.InitializeHTTP(&config)
	if err != nil {
		return nil, err
	}

	provider.jwt = appy_http.NewJwtAuth(APISecret)
	provider.appKey = appy_http.NewApiKeyMiddlewareProvider(APIKey, http.StatusForbidden)

	// Initialize root groups
	provider.Root = provider.Server.Root()
	provider.App = provider.Root.Group("/app")
	provider.Api = provider.Root.Group("/api")
	provider.Public = provider.Api.Group("/")
	provider.Private = provider.Api.Group("/")

	// Setup middleware
	provider.App.Use(provider.appKey.Provide())
	provider.Private.Use(provider.jwt.Authentication())

	return provider, nil
}
