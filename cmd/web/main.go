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

	application := &application{
		errorLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		infoLog:  log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", application.home)
	mux.HandleFunc("/snippet", application.showSnippet)
	mux.HandleFunc("/snippet/create", application.createSnippet)

	// Handle static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: application.errorLog,
		Handler:  mux,
	}

	application.infoLog.Printf("Starting server on %s\n", *addr)
	err := srv.ListenAndServe()
	application.errorLog.Fatal(err)
}
