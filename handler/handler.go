package handler

import (
	"net/http"

	"golang.org/x/net/context"
)

type Handler struct {
}

// ERROR: no usages found for this method even though it is used
func (f *Handler) Method(ctx context.Context, res http.ResponseWriter, req *http.Request) error {
	return nil
}
