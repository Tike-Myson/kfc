package postgresql

import (
	"database/sql"
	"github.com/Tike-Myson/kfc/pkg/models"
	geojson "github.com/paulmach/go.geojson"
)

var filename = "geo.json"

type GeojsonModel struct {
	DB *sql.DB
}

func (m *GeojsonModel) Insert(geom string) error {
	insertStmt := `INSERT INTO "geometries"("id", "geom") values (nextval('geo_sequence'), ST_AsText(ST_GeomFromGeoJSON($1)))`
	_, err := m.DB.Exec(insertStmt, geom)
	if err != nil {
		return err
	}
	return nil
}

func (m *GeojsonModel) Get() (*models.FeatureCollection, error) {
	fc1 := models.NewFeatureCollection()

	rows, err := m.DB.Query(`SELECT "id", ST_AsGeoJSON("geom") FROM "geometries"`)
	if err != nil {
		return fc1, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var geom *geojson.Geometry

		err = rows.Scan(&id, &geom)
		if err != nil {
			return fc1, err
		}
		var fc2 *models.Feature
		fc2.ID = id
		fc2.Geometry = geom
		fc1.Features = append(fc1.Features, fc2)
	}
	return fc1, nil
}

func (m *GeojsonModel) Search() {

}