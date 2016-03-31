package app

import (
	"net/http"

	"github.com/gabrielf/goplugin-usage-bug/handler"
	"golang.org/x/net/context"
)

type App struct {
	SamePkgHandler *Handler
	Handler        *handler.Handler
}

type Handler struct {
}

// usages found for this method
func (f *Handler) Method(ctx context.Context, res http.ResponseWriter, req *http.Request) error {
	return nil
}
