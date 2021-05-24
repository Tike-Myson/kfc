package postgresql

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type GeojsonModel struct {
	DB *sql.DB
}

func (m *GeojsonModel) Insert(name, geom string) error {
	insertStmt := `INSERT INTO "geometries"("name", "geom") values($1, ST_AsText(ST_GeomFromGeoJSON($2)))`
	_, err := m.DB.Exec(insertStmt, name, geom)
	if err != nil {
		return err
	}
	return nil
}

func (m *GeojsonModel) Get() (map[string]string, error) {
	result := make(map[string]string)
	rows, err := m.DB.Query(`SELECT "name", ST_AsGeoJSON("geom") FROM "geometries"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var geom string

		err = rows.Scan(&name, &geom)
		if err != nil {
			return nil, err
		}
		byt := []byte(geom)
		var dat map[string]interface{}
		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}
		fmt.Println(dat)

		result[name] = geom
	}
	return result, nil
}

func (m *GeojsonModel) Search() {

}