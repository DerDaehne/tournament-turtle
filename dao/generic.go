package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo is the data access object for our db itself
type Mongo struct {
	Server   string
	Database string
}

// db is the database to use
var db *mongo.Database

// Connect to a running MongoDB instance
func (dao Mongo) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dao.Server))
	if err != nil {
		return err
	}
	db = client.Database(dao.Database)
	db.ReadPreference()
	return nil
}
