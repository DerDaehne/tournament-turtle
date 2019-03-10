package handles

import (
	"github.com/DerDaehne/tournament-turtle/dao"
	log "github.com/sirupsen/logrus"
)

// Sever struct
var server = dao.Mongo{}

// init initialises the runtime environment
func init() {
	log.Info("Initialize Runtime environment...")
	server.Server = "mongodb://localhost:27017"
	server.Database = "tournamentturtle"

	if err := server.Connect(); err != nil {
		log.Fatal(err)
	}
}
