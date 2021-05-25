package main

import (
	"encoding/json"
	"github.com/Tike-Myson/kfc/pkg/models"
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
