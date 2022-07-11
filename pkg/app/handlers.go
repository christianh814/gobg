package app

import (
	"html/template"
	"net/http"
)

// Set up struct
type Test struct {
	Greeting string
}

// Locate where the templates are
const IndexHtml string = "html/index.tmpl"
const InfoHtml string = "html/info.tmpl"

// appRoot is the HTTP root of the application
func appRoot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(IndexHtml))
	greet := Test{
		Greeting: "Hello, World!",
	}

	// Display index page from template
	tmpl.Execute(w, greet)
}

// appInfo is the route that returns the pod information
func appInfo(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(InfoHtml))

	config := Test{
		Greeting: "This information is for debugging only.",
	}

	// Display index page from template
	tmpl.Execute(w, config)
}
