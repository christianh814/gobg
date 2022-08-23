package app

import (
	"html/template"
	"net/http"
)

// Set up struct
type AppSetting struct {
	Greeting string
	Color    string
}

// Locate where the templates are
const IndexHtml string = "html/index.tmpl"
const InfoHtml string = "html/info.tmpl"

// appRoot is the HTTP root of the application
func appRoot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(IndexHtml))
	greet := AppSetting{
		Greeting: "Purpleiris",
		Color:    "purple",
	}

	// Display index page from template
	tmpl.Execute(w, greet)
}

// appInfo is the route that returns the pod information
func appInfo(w http.ResponseWriter, r *http.Request) {
	// this needs to be updated when the /info route is implemented
	tmpl := template.Must(template.ParseFiles(InfoHtml))

	greet := AppSetting{
		Greeting: "This information is for debugging only.",
	}

	// Display index page from template
	tmpl.Execute(w, greet)
}
