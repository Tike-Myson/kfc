package postgresql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
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

func (m *GeojsonModel) Get() error {
	rows, err := m.DB.Query(`SELECT "id", ST_AsGeoJSON("geom") FROM "geometries"`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		//var geom *geojson.Geometry
		var geom string

		err = rows.Scan(&id, &geom)
		if err != nil {
			return err
		}
		fmt.Println(id, geom)
	}
	return nil
}

func writeJsonToFile(data []byte) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *GeojsonModel) Search() {

}