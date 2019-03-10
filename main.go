// Package main contains the main routine of the turnament tourtle application
package main

import (
	"fmt"
	"net/http"

	"github.com/DerDaehne/tournament-turtle/handles"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// main creates a new mux Router and starts listening on a network port
func main() {

	// create new router
	router := mux.NewRouter()

	// ignore this - i need this so that gofmt thinks the package "fmt" is in use and won't delete it from my import list
	fmt.Printf("Hello\n")

	// set handler functions for players
	router.HandleFunc("/players", handles.AllPlayersEndPoint).Methods("GET")
	router.HandleFunc("/players", handles.CreatePlayerEndPoint).Methods("POST")
	router.HandleFunc("/players", handles.UpdatePlayerEndPoint).Methods("PUT")
	router.HandleFunc("/players", handles.DeletePlayerEndPoint).Methods("DELETE")
	router.HandleFunc("/players/{id}", handles.FindPlayerByIDEndPoint).Methods("GET")

	// set handler functions for teams
	router.HandleFunc("/teams", handles.AllTeamsEndPoint).Methods("GET")
	router.HandleFunc("/teams", handles.CreateTeamEndPoint).Methods("POST")
	router.HandleFunc("/teams", handles.UpdateTeamEndPoint).Methods("PUT")
	router.HandleFunc("/teams", handles.DeleteTeamEndPoint).Methods("DELETE")
	router.HandleFunc("/teams/{id}", handles.FindTeamByIDEndPoint).Methods("GET")

	// start listening on port 8080
	// ports is currently hard coded and will be configurable later on
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
