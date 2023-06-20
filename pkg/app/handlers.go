package app

import (
	"html/template"
	"net/http"
	"os"
)

// Set up struct
type AppSetting struct {
	Greeting string
	Color    string
}

// Info type
type Info struct {
	Alert  string
	Status string
	Info   map[string]string
	Msg    string
}

// Locate where the templates are
const IndexHtml string = "html/index.tmpl"
const InfoHtml string = "html/info.tmpl"

// appRoot is the HTTP root of the application
func appRoot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(IndexHtml))
	greet := AppSetting{
		Greeting: "Just a Blue Box",
		Color:    "blue",
	}

	// Display index page from template
	tmpl.Execute(w, greet)
}

// appInfo is the route that returns the pod information
func appInfo(w http.ResponseWriter, r *http.Request) {
	// this needs to be updated when the /info route is implemented
	tmpl := template.Must(template.ParseFiles(InfoHtml))
	info := map[string]string{
		"Hostname": os.Getenv("HOSTNAME"),
		"App IP":   os.Getenv("GOBG_SERVICE_HOST"),
		"App Port": os.Getenv("GOBG_PORT_8080_TCP_PORT"),
	}

	config := Info{
		Info:   info,
		Status: "App Info:",
		Alert:  "alert-primary",
		Msg:    "This information is for debugging only.",
	}

	// Display index page from template
	tmpl.Execute(w, config)
}
