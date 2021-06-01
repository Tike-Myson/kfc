package postgresql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	geojson "github.com/paulmach/go.geojson"
)

type GeojsonModel struct {
	DB *sql.DB
}

func (m *GeojsonModel) Insert(fc *geojson.FeatureCollection) error {
	for _, v := range fc.Features {
		e, err := json.Marshal(v.Geometry)
		if err != nil {
			return err
		}
		insertStmt := `INSERT INTO "geometries"("id", "geom") values (nextval('geo_sequence'), ST_AsText(ST_GeomFromGeoJSON($1)))`
		_, err = m.DB.Exec(insertStmt, string(e))
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *GeojsonModel) Get() (*geojson.FeatureCollection, error) {
	fc := geojson.NewFeatureCollection()
	rows, err := m.DB.Query(`SELECT "id", ST_AsGeoJSON("geom") FROM "geometries"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var geom *geojson.Geometry
		err = rows.Scan(&id, &geom)
		if err != nil {
			return nil, err
		}
		f := geojson.NewFeature(geom)
		f.ID = id
		fc = fc.AddFeature(f)
	}
	return fc, nil
}

func (m *GeojsonModel) Search(fc *geojson.FeatureCollection) (*geojson.FeatureCollection, error){
	fcResult := geojson.NewFeatureCollection()
	fmt.Println(fc)
	for _, v := range fc.Features {
		e, err := json.Marshal(v.Geometry)
		if err != nil {
			fmt.Println(string(e))
			rows, err := m.DB.Query(`SELECT "id", ST_AsGeoJSON("geom") FROM geometries WHERE ST_Intersects(geom, $1)`, string(e))
			if err != nil {
				return nil, err
			}
			defer rows.Close()

			for rows.Next() {
				var id int
				var geom *geojson.Geometry
				err = rows.Scan(&id, &geom)
				if err != nil {
					return nil, err
				}
				f := geojson.NewFeature(geom)
				f.ID = id
				fcResult = fcResult.AddFeature(f)
			}
		}
	}
	return fcResult, nil
}