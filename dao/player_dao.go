package dao

import (
	"context"
	"log"
	"time"

	"github.com/DerDaehne/tournament-turtle/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// PlayerDAO is the data access object for our players db
type PlayerDAO struct {
	Server string
}

const (
	// COLLECTION is the name of the mongodb collection to use
	COLLECTION = "players"
)

// client is the mongoDB Client
var client mongo.Client

// Connect to a running MongoDB instance
func (dao *PlayerDAO) Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dao.Server))
	if err != nil {
		log.Fatal(err)
	}
	client.Ping(ctx, readpref.Primary())
}

// Insert a new Entry into our Collection
func (dao *PlayerDAO) Insert(player models.Player) {

}
