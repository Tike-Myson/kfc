package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := mux.NewRouter().StrictSlash(true)

	return secureHeaders(app.recoverPanic(app.logRequest(mux)))
}