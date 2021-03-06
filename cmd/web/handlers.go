package main

import (
	"encoding/json"
	"fmt"
	"github.com/alecthomas/template"
	_ "github.com/aws/aws-lambda-go/events"
	"github.com/gorilla/mux"
	geojson "github.com/paulmach/go.geojson"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		app.serverError(w, err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) returnAPI(w http.ResponseWriter, r *http.Request){
	content := readJsonFile()
	fc2:= geojson.NewFeatureCollection()
	json.Unmarshal(content, fc2)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	json.NewEncoder(w).Encode(fc2)
}

func (app *application) returnSingleAPI(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	content := readJsonFile()
	fc2:= geojson.NewFeatureCollection()
	json.Unmarshal(content, fc2)
	for i, v := range fc2.Features {
		v.ID = i
		id := fmt.Sprintf("%v", v.ID)
		if key == id {
			fc3 := geojson.NewFeature(v.Geometry)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Origin")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Accept", "application/json")
			json.NewEncoder(w).Encode(fc3)
		}
	}
}

func (app *application) API(w http.ResponseWriter, r *http.Request){
	fc := geojson.NewFeatureCollection()
	err := json.NewDecoder(r.Body).Decode(&fc)
	if err != nil {
		app.serverError(w, err)
	}
	fmt.Fprintf(w, "%v", fc)
	fcResult, err := app.geometries.SmartSearch(fc)
	if err != nil {
		app.serverError(w, err)
	}
	writeJson(fcResult)

}
