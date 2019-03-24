package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"

	"github.com/DerDaehne/tournament-turtle/models"
)

// PlayerDAO is the data access object for our players db collection
type PlayerDAO struct {
}

const (
	// PCOLLECTION is the name of the mongodb collection to use
	PCOLLECTION = "players"
)

// FindAll returns all returns all document in the given collection
func (dao *PlayerDAO) FindAll() ([]models.Player, error) {
	var players []models.Player
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.Collection(PCOLLECTION).Find(ctx, bson.D{})
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Collection(PCOLLECTION).InsertOne(ctx, bson.D{
		{"firstname", player.FirstName},
		{"lastname", player.LastName},
		{"nickname", player.NickName},
		{"skilllevel", player.SkillLevel},
	})
	return err
}

// FindByID searches for an Player by an given id
func (dao *PlayerDAO) FindByID(playerID string) (models.Player, error) {
	var player models.Player

	id, errconvert := primitive.ObjectIDFromHex(playerID)
	if errconvert != nil {
		return player, errconvert
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}

	err := db.Collection(PCOLLECTION).FindOne(ctx, filter).Decode(&player)
	if err != nil {
		return player, err
	}
	return player, nil
}

// Delete a player entry
func (dao *PlayerDAO) Delete(player models.Player) error {
	id, errconvert := primitive.ObjectIDFromHex(player.ID)
	if errconvert != nil {
		return errconvert
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}

	_, err := db.Collection(PCOLLECTION).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// Update a entry in the database
func (dao *PlayerDAO) Update(player models.Player) error {
	id, errconvert := primitive.ObjectIDFromHex(player.ID)
	if errconvert != nil {
		return errconvert
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}

	_, err := db.Collection(PCOLLECTION).UpdateOne(ctx, filter, bson.D{
		{"$set",
			bson.D{
				{"firstname", player.FirstName},
				{"lastname", player.LastName},
				{"nickname", player.NickName},
				{"skilllevel", player.SkillLevel},
			}},
	})
	if err != nil {
		return err
	}
	return nil
}
