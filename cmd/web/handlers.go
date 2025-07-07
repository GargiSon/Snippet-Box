package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Change the signature of the handler so it is defined as a method against *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Because the home handler is now a method against application struct it can access its fields, including the structured logger. We'll use this to create a log entry level containing the error message, also including the request method and URI as attributes to assisst with debugging.
		// app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		//These 2 lines in one
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		// These 2 lines in one
		app.serverError(w, r, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
