package main

import (
	"encoding/json"
	"fmt"
	"github.com/Tike-Myson/kfc/pkg/models"
	geojson "github.com/paulmach/go.geojson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

var filename = "test.json"

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

func writeJsonToFile(id string, geom *geojson.Geometry) error {
	content := readJsonFile()
	fc1 := models.NewFeatureCollection()
	err := json.Unmarshal(content, fc1)
	if err != nil {
		return err
	}
	var fc2 *models.Feature
	fc2.ID = id
	fc2.Geometry = geom
	fc1.Features = append(fc1.Features, fc2)
	data, err := json.Marshal(fc1)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
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
