// Package main contains the main routine of the turnament tourtle application
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// logRequestInfo logs mostly useful information about received requests
func logRequestInfo(request *http.Request) {
	// log at "info" level
	log.WithFields(log.Fields{
		"RemoteAddr": request.RemoteAddr,
		"Method":     request.Method,
		"RequestURI": request.RequestURI,
	}).Info()
}

// AllPlayersEndPoint returns all Players currently in the database
func AllPlayersEndPoint(writer http.ResponseWriter, request *http.Request) {
	logRequestInfo(request)
	fmt.Fprintln(writer, "not implemented yet!")
}

// CreatePlayerEndPoint will create a new Player entry in the database
func CreatePlayerEndPoint(writer http.ResponseWriter, request *http.Request) {
	logRequestInfo(request)
	fmt.Fprintln(writer, "not implemented yet!")
}

// UpdatePlayerEndPoint will update a Player entry in the database
func UpdatePlayerEndPoint(writer http.ResponseWriter, request *http.Request) {
	logRequestInfo(request)
	fmt.Fprintln(writer, "not implemented yet!")
}

// DeletePlayerEndPoint will drop a Player entry
func DeletePlayerEndPoint(writer http.ResponseWriter, request *http.Request) {
	logRequestInfo(request)
	fmt.Fprintln(writer, "not implemented yet!")
}

// FindPlayerEndPoint will find a Player's entry
func FindPlayerEndPoint(writer http.ResponseWriter, request *http.Request) {
	logRequestInfo(request)
	fmt.Fprintln(writer, "not implemented yet!")
}

// main creates a new mux Router and starts listening on a network port
func main() {
	// create new router
	router := mux.NewRouter()

	// ignore this - i need this so that gofmt thinks the package "fmt" is in use and won't delete it from my import list
	fmt.Printf("Hello")

	// set handler functions
	router.HandleFunc("/players", AllPlayersEndPoint).Methods("GET")
	router.HandleFunc("/players", CreatePlayerEndPoint).Methods("POST")
	router.HandleFunc("/players", UpdatePlayerEndPoint).Methods("PUT")
	router.HandleFunc("/players", DeletePlayerEndPoint).Methods("DELETE")
	router.HandleFunc("/players", FindPlayerEndPoint).Methods("GET")

	// start listening on port 8080
	// ports is currently hard coded and will be configurable later on
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
