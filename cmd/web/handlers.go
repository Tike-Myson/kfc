package main

import (
	"encoding/json"
	geojson "github.com/paulmach/go.geojson"
	"net/http"
)

func (app *application) returnScoreboard(w http.ResponseWriter, r *http.Request){
	content := readJsonFile()
	fc2, err := geojson.UnmarshalGeometry(content)
	if err != nil {
		app.serverError(w, err)
	}
	json.Unmarshal(content, fc2)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	json.NewEncoder(w).Encode(fc2)
}
