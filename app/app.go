package app

import "github.com/gabrielf/go-sandbox/handler"

type App struct {
	SamePkgHandler *Handler
	Handler        *handler.Handler
}

type Handler struct {
}

// usages found for this method
func (f *Handler) Method() {}
