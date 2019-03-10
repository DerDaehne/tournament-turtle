package handles

import (
	"encoding/json"
	"net/http"

	"github.com/DerDaehne/tournament-turtle/dao"
	"github.com/DerDaehne/tournament-turtle/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// create a new dao
var playerDAO = dao.PlayerDAO{}

// AllPlayersEndPoint returns all Players currently in the database
func AllPlayersEndPoint(writer http.ResponseWriter, request *http.Request) {
	logRequestInfo(request)

	players, err := playerDAO.FindAll()
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}

	respondWithJSON(writer, http.StatusOK, players)
}

// CreatePlayerEndPoint will create a new Player entry in the database
func CreatePlayerEndPoint(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	logRequestInfo(request)

	var player models.Player

	if err := json.NewDecoder(request.Body).Decode(&player); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}
	if err := playerDAO.Insert(player); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}

	respondWithJSON(writer, http.StatusOK, map[string]string{"result": "success"})
}

// UpdatePlayerEndPoint will update a Player entry in the database
func UpdatePlayerEndPoint(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	logRequestInfo(request)

	var player models.Player

	if err := json.NewDecoder(request.Body).Decode(&player); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}
	if err := playerDAO.Update(player); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}

	respondWithJSON(writer, http.StatusOK, map[string]string{"result": "success"})
}

// DeletePlayerEndPoint will drop a Player entry
func DeletePlayerEndPoint(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	logRequestInfo(request)

	var player models.Player

	if err := json.NewDecoder(request.Body).Decode(&player); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}
	if err := playerDAO.Delete(player); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}

	respondWithJSON(writer, http.StatusOK, map[string]string{"result": "success"})
}

// FindPlayerByIDEndPoint will find a Player's entry
func FindPlayerByIDEndPoint(writer http.ResponseWriter, request *http.Request) {
	logRequestInfo(request)

	parameters := mux.Vars(request)
	log.Info(parameters["id"])
	player, err := playerDAO.FindByID(parameters["id"])
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}

	respondWithJSON(writer, http.StatusOK, player)
}
