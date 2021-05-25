package main

import (
	"encoding/json"
	"fmt"
	"github.com/Tike-Myson/kfc/pkg/models"
	geojson "github.com/paulmach/go.geojson"
	"net/http"
)

func (app *application) returnAPI(w http.ResponseWriter, r *http.Request){
	content := readJsonFile()
	fc2, _ := geojson.UnmarshalGeometry(content)
	var test models.Geometries
	json.Unmarshal(content, test)
	fmt.Println(test)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	json.NewEncoder(w).Encode(fc2)
}
