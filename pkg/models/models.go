package models

import geojson "github.com/paulmach/go.geojson"

type Feature struct {
	ID          interface{}            `json:"id,omitempty"`
	BoundingBox []float64              `json:"bbox,omitempty"`
	Geometry    *geojson.Geometry      `json:"geometry"`
	CRS         map[string]interface{} `json:"crs,omitempty"` // Coordinate Reference System Objects are not currently supported
}

type FeatureCollection struct {
	BoundingBox []float64              `json:"bbox,omitempty"`
	Features    []*Feature             `json:"features"`
	CRS         map[string]interface{} `json:"crs,omitempty"` // Coordinate Reference System Objects are not currently supported
}

// NewFeatureCollection creates and initializes a new feature collection.
func NewFeatureCollection() *FeatureCollection {
	return &FeatureCollection{
		Features: make([]*Feature, 0),
	}
}