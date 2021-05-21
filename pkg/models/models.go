package models

import geojson "github.com/paulmach/go.geojson"

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Time  string `json:"time"`
}
var Scoreboard2 []geojson.Feature
var Scoreboard []Player