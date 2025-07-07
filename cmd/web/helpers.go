package main

import (
	"net/http"
	"runtime/debug"
)

// The serverError helper writes a log entry at Error level (including the request method and URI as attributes), then sends a generic 500 internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()

		//Use debug.Stack() to get the stack trace. This returns a byte slice, which we need to convert so that it's readable in the log entry.
		trace = string(debug.Stack())
	)
	// And included the trace in the log entry
	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// This function sends a specific status code and corresponding description
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
