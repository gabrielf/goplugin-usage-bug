package main

import (
	"github.com/andviro/noodle"
	"github.com/gabrielf/go-sandbox/app"
	"github.com/gorilla/mux"
)

func main() {
	app := app.App{}

	middlewareChain := noodle.New()

	router := mux.NewRouter()
	// implementation found for both of these methods
	router.Handle("/path", middlewareChain.Then(app.SamePkgHandler.Method))
	router.Handle("/path", middlewareChain.Then(app.Handler.Method))
}
