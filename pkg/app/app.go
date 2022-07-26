package app

import (
	"net/http"
	"os"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// static port for now
const HttpPort string = "8080"

func Start() {
	//create/register a new request multiplexer
	router := mux.NewRouter()

	// Set up path to static assets
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("/app/html/assets/"))))

	// register the / route.
	router.Handle("/", gh.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(appRoot)))

	//router.Handle("/info", gh.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(appInfo)))

	// try to start the app and log output.
	log.Info("Starting server on port " + HttpPort)
	log.Fatal(http.ListenAndServe(":"+HttpPort, router))
}
