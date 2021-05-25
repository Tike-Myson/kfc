package main

import (
	_ "github.com/Tike-Myson/kfc/docs"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/", app.returnAPI).Methods("GET")
	mux.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	return secureHeaders(app.recoverPanic(app.logRequest(mux)))
}