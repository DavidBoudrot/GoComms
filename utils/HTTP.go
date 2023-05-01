package utils

//HTTPStartup will be called in main to start up the HTTP server

import (
	"log"
	"net/http"
)

var HTMLport = "8080"

func HTTPStartup() {
	// This will be the main function that will start up the HTTP server
	// It will also be the function that will call the other functions
	// that will handle the HTTP requests
	log.Println("Starting HTTP server on port", HTMLport)
	http.HandleFunc("/", makeHandler(indexHandler))

}

func makeHandler(handler interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.(func(http.ResponseWriter, *http.Request))(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
