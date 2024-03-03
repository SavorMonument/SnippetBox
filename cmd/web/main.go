package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"zzz1.org/snippetbox/internal/models"
)

// Application wide dependecies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	addr := os.Getenv("ADDRESS")
	dsn := os.Getenv("DSN")

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting server on %s\n", addr)
	err = srv.ListenAndServe()
	app.errorLog.Fatal(err)
}
