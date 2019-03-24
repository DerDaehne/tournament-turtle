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
var teamDAO = dao.TeamDAO{}

// AllTeamsEndPoint returns all Teams currently in the database
func AllTeamsEndPoint(writer http.ResponseWriter, request *http.Request) {
	logRequestInfo(request)

	teams, err := teamDAO.FindAll()
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}

	respondWithJSON(writer, http.StatusOK, teams)
}

// CreateTeamEndPoint will create a new Team entry in the database
func CreateTeamEndPoint(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	logRequestInfo(request)

	var team models.Team

	if err := json.NewDecoder(request.Body).Decode(&team); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}
	if err := teamDAO.Insert(team); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}

	respondWithJSON(writer, http.StatusOK, map[string]string{"result": "success"})
}

// UpdateTeamEndPoint will update a Team entry in the database
func UpdateTeamEndPoint(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	logRequestInfo(request)

	var team models.Team

	if err := json.NewDecoder(request.Body).Decode(&team); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}
	if err := teamDAO.Update(team); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}

	respondWithJSON(writer, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteTeamEndPoint will drop a Team entry
func DeleteTeamEndPoint(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	logRequestInfo(request)

	var team models.Team

	if err := json.NewDecoder(request.Body).Decode(&team); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}
	if err := teamDAO.Delete(team); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}

	respondWithJSON(writer, http.StatusOK, map[string]string{"result": "success"})
}

// FindTeamByIDEndPoint will find a Team's entry
func FindTeamByIDEndPoint(writer http.ResponseWriter, request *http.Request) {
	logRequestInfo(request)

	parameters := mux.Vars(request)
	log.Info(parameters["id"])
	team, err := teamDAO.FindByID(parameters["id"])
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		log.Error(err)
		return
	}

	respondWithJSON(writer, http.StatusOK, team)
}
