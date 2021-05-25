package main

import (
	"encoding/json"
	"fmt"
	"github.com/Tike-Myson/kfc/pkg/models"
	"github.com/gorilla/mux"
	geojson "github.com/paulmach/go.geojson"
	"net/http"
)

func (app *application) returnAPI(w http.ResponseWriter, r *http.Request){
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
	fmt.Println(key)
	content := readJsonFile()
	fc2:= models.NewFeatureCollection()
	json.Unmarshal(content, fc2)
	for _, v := range fc2.Features {
		id := fmt.Sprintf("%v", v.ID)
		if key == id {
			fc3 := geojson.NewFeature(v.Geometry)
			fmt.Println(fc3)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Origin")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Accept", "application/json")
			json.NewEncoder(w).Encode(fc3)
		}
	}
}

