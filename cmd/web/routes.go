package main

import (
	_ "github.com/Tike-Myson/kfc/docs"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/", app.home).Methods("GET")
	mux.HandleFunc("/api", app.returnAPI).Methods("GET")
	mux.HandleFunc("/api", app.API).Methods("POST")
	mux.HandleFunc("/api/{id}", app.returnSingleAPI).Methods("GET")
	mux.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))

	return secureHeaders(app.recoverPanic(app.logRequest(mux)))
}