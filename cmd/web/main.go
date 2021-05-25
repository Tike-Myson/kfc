package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Tike-Myson/kfc/pkg/models/postgresql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	//geometries interface{
	//	Insert(string, string) error
	//	Get() (*models.Geometries, error)
	//	Search()
	//}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "nurtilek"
	password = "nm2000kz"
	dbname   = "postgis_db"
)

// @title KFC Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, Green+"INFO\t"+Reset, log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, Red+"ERROR\t"+Reset, log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}



	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var g postgresql.GeojsonModel
	g.DB = db
	//err = g.Insert(string(readJsonFile()))
	//if err != nil {
	//	errorLog.Panicln(err)
	//}
	err = g.Get()
	if err != nil {
		errorLog.Panicln(err)
	}
	infoLog.Printf("Server run on http://127.0.0.1%s\n", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
