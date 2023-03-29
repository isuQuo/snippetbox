package main

import "net/http"

func (app *application) routes() http.Handler {
	// Create a new ServeMux and register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static/" directory.
	//fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Use the mux.Handle() function to register the file server as the handler for all URL paths that start with "/static/".
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return secureHeaders(mux)
}
