package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
func (dao *PlayerDAO) Connect() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dao.Server))
	if err != nil {
		return err
	}
	db = client.Database(dao.Database)
	db.ReadPreference()
	return nil
}

// FindAll returns all returns all document in the given collection
func (dao *PlayerDAO) FindAll() ([]models.Player, error) {
	var players []models.Player
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := db.Collection(COLLECTION).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		documents := &bson.D{}
		if err := cursor.Decode(documents); err != nil {
			log.Fatal(err)
		}

		m := documents.Map()
		player := models.Player{
			ID:         m["_id"].(primitive.ObjectID).Hex(),
			LastName:   m["lastname"].(string),
			FirstName:  m["firstname"].(string),
			NickName:   m["nickname"].(string),
			SkillLevel: int(m["skilllevel"].(int32)),
		}
		players = append(players, player)
	}
	return players, nil
}

// Insert a new Entry into our Collection
func (dao *PlayerDAO) Insert(player models.Player) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.Collection(COLLECTION).InsertOne(ctx, bson.D{
		{"firstname", player.FirstName},
		{"lastname", player.LastName},
		{"nickname", player.NickName},
		{"skilllevel", player.SkillLevel},
	})
	return err
}
