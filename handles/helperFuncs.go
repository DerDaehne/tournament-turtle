package handles

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

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

// logRequestInfo logs mostly useful information about received requests
func logRequestInfo(request *http.Request) {
	// log at "info" level
	log.WithFields(log.Fields{
		"RemoteAddr": request.RemoteAddr,
		"Method":     request.Method,
		"RequestURI": request.RequestURI,
	}).Info()
}
