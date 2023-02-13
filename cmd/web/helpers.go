package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

// The serverError helper writes an error message and stack trace to the errorLog,
// then sneds a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad Request"
// when there's a problem with the request that the user sent.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to
// the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// Create an addDefaultData helper. This takes a pointer to a templateData
// struct and then returns the pointer.
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	td.AuthenticatedUser = app.authenticatedUser(r)
	td.CurrentYear = time.Now().Year()
	td.Flash = app.session.PopString(r, "flash")
	return td
}

// The authenticatedUser method returns the ID of the current user from the
// session, or zero if the request is from an authenticated user.
func (app *application) authenticatedUser(r *http.Request) int {
	return app.session.GetInt(r, "userID")
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	// Retrieve the appropriate template set from the cache based on the page name
	// (like 'home.page.tmpl'). If no entry exists in the cache with the
	// provided name, call the serverError helper method that we made earlier.
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exists", name))
		return
	}

	// Create a new buffer
	buf := new(bytes.Buffer)

	// Write the template to buffer instead of straight to the http.ResponseWriter
	// To catch a runtime error
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
	}

	// Write the content to http.ResponseWriter as it is safe to do so
	buf.WriteTo(w)
}
