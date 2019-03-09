// Package main contains the main routine of the turnament tourtle application
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DerDaehne/tournament-turtle/dao"
	"github.com/DerDaehne/tournament-turtle/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var playerDAO = dao.PlayerDAO{}

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
	defer request.Body.Close()
	logRequestInfo(request)

	var player models.Player

	if err := json.NewDecoder(request.Body).Decode(&player); err != nil {
		log.Error("json decode error: " + err.Error())
		return
	}
	if err := playerDAO.Insert(player); err != nil {
		log.Fatal(err)
	}

	respondWithJSON(writer, http.StatusOK, map[string]string{"result": "success"})
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

// respondWithError returns a http status code
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

// respondWithJson wraps a given payload into json
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// init initialises the runtime environment
func init() {
	log.Info("Initialize Runtime environment...")
	playerDAO.Server = "mongodb://localhost:27017"
	playerDAO.Database = "tournamentturtle"

	if err := playerDAO.Connect(); err != nil {
		log.Fatal(err)
	}
}

// main creates a new mux Router and starts listening on a network port
func main() {

	// create new router
	router := mux.NewRouter()

	// ignore this - i need this so that gofmt thinks the package "fmt" is in use and won't delete it from my import list
	fmt.Printf("Hello\n")

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
