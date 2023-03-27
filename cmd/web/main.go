package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Create a new ServeMux and register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static/" directory.
	//fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Use the mux.Handle() function to register the file server as the handler for all URL paths that start with "/static/".
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr:     "127.0.0.1:8080",
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Println("Starting server on :8080")
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
