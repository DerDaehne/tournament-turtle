package dao

import (
	"context"
	"time"

	"github.com/DerDaehne/tournament-turtle/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TeamDAO is the data access object for our teams db collection
type TeamDAO struct{}

const (
	// TCOLLECTION is the name of the mongodb collection to use
	TCOLLECTION = "teams"
)

// FindAll returns all returns all document in the given collection
func (dao *TeamDAO) FindAll() ([]models.Team, error) {
	var teams []models.Team
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.Collection(TCOLLECTION).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		document := &bson.D{}
		if err := cursor.Decode(document); err != nil {
			log.Fatal(err)
		}

		teamP := &models.Team{}

		b, err := bson.Marshal(document)
		if err != nil {
			return nil, err
		}

		errM := bson.Unmarshal(b, teamP)
		if errM != nil {
			return nil, err
		}

		teams = append(teams, *teamP)
	}
	return teams, nil
}

// Insert a new Entry into our Collection
func (dao *TeamDAO) Insert(team models.Team) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Collection(TCOLLECTION).InsertOne(ctx, bson.D{
		{"name", team.Name},
		{"member", team.Member},
	})
	return err
}

// FindByID searches for an Player by an given id
func (dao *TeamDAO) FindByID(teamID string) (models.Team, error) {
	var team models.Team

	id, errconvert := primitive.ObjectIDFromHex(teamID)
	if errconvert != nil {
		return team, errconvert
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}

	err := db.Collection(TCOLLECTION).FindOne(ctx, filter).Decode(&team)
	if err != nil {
		return team, err
	}
	return team, nil
}

// Delete a team entry
func (dao *TeamDAO) Delete(team models.Team) error {
	id, errconvert := primitive.ObjectIDFromHex(team.ID)
	if errconvert != nil {
		return errconvert
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}

	_, err := db.Collection(TCOLLECTION).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// Update a entry in the database
func (dao *TeamDAO) Update(team models.Team) error {
	id, errconvert := primitive.ObjectIDFromHex(team.ID)
	if errconvert != nil {
		return errconvert
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}

	_, err := db.Collection(TCOLLECTION).UpdateOne(ctx, filter, bson.D{
		{"$set",
			bson.D{
				{"name", team.Name},
				{"member", team.Member},
			}},
	})
	if err != nil {
		return err
	}
	return nil
}
