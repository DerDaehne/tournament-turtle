// Package main contains the main routine of the turnament tourtle application
package main

import (
	"fmt"
	"net/http"

	"github.com/DerDaehne/tournament-turtle/handles"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware", r.Method)

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			return
		}

		next.ServeHTTP(w, r)
		log.Println("Executing middleware again")
	})
}

// main creates a new mux Router and starts listening on a network port
func main() {

	// create new router
	router := mux.NewRouter()

	// ignore this - i need this so that gofmt thinks the package "fmt" is in use and won't delete it from my import list
	fmt.Printf("Hello\n")

	router.HandleFunc("/players", handles.AllPlayersEndPoint).Methods("GET")
	router.HandleFunc("/players", handles.CreatePlayerEndPoint).Methods("POST", "OPTIONS")
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
	//router.Use(mux.CORSMethodMiddleware(router))
	//if err := http.ListenAndServe(":8080", router); err != nil {
	if err := http.ListenAndServe(":8080", corsMiddleware(router)); err != nil {
		log.Fatal(err)
	}
}
