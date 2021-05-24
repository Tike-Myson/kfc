package main

import (
	"fmt"
	"github.com/Tike-Myson/kfc/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"sort"
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

func writeJsonToFile(data []byte) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func sortData() {
	sort.SliceStable(models.Scoreboard, func(i, j int) bool {
		return models.Scoreboard[i].Score > models.Scoreboard[j].Score
	})
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
