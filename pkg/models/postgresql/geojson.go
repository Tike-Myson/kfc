package postgresql

import (
	"database/sql"
	"github.com/Tike-Myson/kfc/pkg/models"
	geojson "github.com/paulmach/go.geojson"
	//"github.com/Tike-Myson/kfc/pkg/models"
	//geojson "github.com/paulmach/go.geojson"
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
	fc := models.NewFeatureCollection()
	rows, err := m.DB.Query(`SELECT "id", ST_AsGeoJSON("geom") FROM "geometries"`)
	if err != nil {
		return fc, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var geom *geojson.Geometry
		err = rows.Scan(&id, &geom)
		if err != nil {
			return fc, err
		}
		f := models.NewFeature(geom)
		f.ID = id
		fc.AddFeature(f)
	}
	return fc, nil
}

func (m *GeojsonModel) Search() {

}