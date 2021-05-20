package models

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Time  string `json:"time"`
}

var Scoreboard []Player