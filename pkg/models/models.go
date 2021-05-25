package models

import geojson "github.com/paulmach/go.geojson"

type Geometries struct {
	id int `json:"id"`
	Geometry *geojson.Geometry `json:"geometry"`
}
