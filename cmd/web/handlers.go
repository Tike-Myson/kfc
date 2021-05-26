package main

import (
	"encoding/json"
	"fmt"
	"github.com/Tike-Myson/kfc/pkg/models"
	"github.com/gorilla/mux"
	geojson "github.com/paulmach/go.geojson"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (app *application) returnAPI(w http.ResponseWriter, r *http.Request){
	fc, err := app.geometries.Get()
	fmt.Println(fc)
	if err != nil {
		app.serverError(w, err)
	}
	for _, v := range fc.Features {
		str := fmt.Sprintf("%v", v.ID)
		id, err := strconv.Atoi(str)
		if err != nil {
			app.errorLog.Println(err)
		}
		writeJsonToFile(id, v.Geometry)
	}
	content := readJsonFile()
	fc2:= models.NewFeatureCollection()
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
	fc2:= models.NewFeatureCollection()
	json.Unmarshal(content, fc2)
	for _, v := range fc2.Features {
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
	body, err := ioutil.ReadAll(r.Body)
	str := sanitizeGeomJson(body)
	err = app.geometries.Insert(str)
	if err != nil {
		app.serverError(w, err)
	}
}

