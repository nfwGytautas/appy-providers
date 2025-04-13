package providers_http

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type ErrorMapper struct {
}

type ErrorResponse map[string]any

func (em *ErrorMapper) Map(ctx context.Context, err error) (int, any) {
	return http.StatusInternalServerError, ErrorResponse{
		"error": errors.WithStack(err).Error(),
		"toast": "An error occurred while processing your request",
	}
}
