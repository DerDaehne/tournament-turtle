package models

import "go.mongodb.org/mongo-driver/bson"

// Player represents a Person who want's to play
type Player struct {
	ID         bson.D
	LastName   string
	FirstName  string
	NickName   string
	SkillLevel int
}
