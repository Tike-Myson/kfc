package models

import geojson "github.com/paulmach/go.geojson"

type Geometries struct {
	Id int
	Geom *geojson.Geometry
}