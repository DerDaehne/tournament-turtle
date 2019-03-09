package dao

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/DerDaehne/tournament-turtle/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PlayerDAO is the data access object for our players db
type PlayerDAO struct {
	Server   string
	Database string
}

const (
	// COLLECTION is the name of the mongodb collection to use
	COLLECTION = "players"
)

// db is the database to use
var db *mongo.Database

// Connect to a running MongoDB instance
func (dao *PlayerDAO) Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dao.Server))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(dao.Database)
	db.ReadPreference()
}

// Insert a new Entry into our Collection
func (dao *PlayerDAO) Insert(player models.Player) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.Collection(COLLECTION).InsertOne(ctx, player)
	return err
}
