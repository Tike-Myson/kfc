package models

type Geometries struct {
	Name string `json:"name"`
	Geometry Geometry `json:"geom"`
}

type Geometry struct {
	Type string `json:"type"`
	Geo string    `json:"coordinates"`
}
