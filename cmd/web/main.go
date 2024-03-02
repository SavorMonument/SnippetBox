package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Application wide dependecies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := &application{
		errorLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		infoLog:  log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting server on %s\n", *addr)
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)
}
