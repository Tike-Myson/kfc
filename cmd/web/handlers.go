package main

import (
	"encoding/json"
	"net/http"
	"github.com/Tike-Myson/kfc/pkg/models"
)

func (app *application) returnScoreboard(w http.ResponseWriter, r *http.Request){
	content := readJsonFile()
	json.Unmarshal(content, &models.Scoreboard)
	sortData()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	json.NewEncoder(w).Encode(models.Scoreboard)
}
