package main

import (
	"encoding/json"
	"fmt"
	geojson "github.com/paulmach/go.geojson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

var filename = "geo.json"

func readJsonFile() []byte {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fileName, err := os.Create(filename)
		if err != nil {
			log.Fatal(err.Error())
		}
		fileName.Close()
	}

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err.Error())
	}
	return content
}

func writeJson(fc *geojson.FeatureCollection) error {
	data, err := json.Marshal(fc)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("geo.json", data, 0644)
	return nil
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}