package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// Create a new httprouter instance.
	router := httprouter.New()

	// Wrap the http.NotFound() function in a http.HandlerFunc so that it
	// returns our own custom 404 Not Found response.
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	//fileServer := http.FileServer(http.Dir("./ui/static/"))
	//router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave, app.noSurf, app.authenticate)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.signupUserForm))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.signupUser))
	router.Handler(http.MethodGet, "/user/signin", dynamic.ThenFunc(app.signinUserForm))
	router.Handler(http.MethodPost, "/user/signin", dynamic.ThenFunc(app.signinUser))

	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreateForm))
	router.Handler(http.MethodPost, "/snippet/create", protected.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/user/signout", protected.ThenFunc(app.signoutUser))

	// Create a middleware chain containing our 'standard' middleware
	// which will be used by every request our application receives.
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
