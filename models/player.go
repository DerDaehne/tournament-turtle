package models

// Player represents a Person who want's to play
type Player struct {
	LastName   string `bson:"lastname" json:"lastname"`
	FirstName  string `bson:"firstname" json:"firstname"`
	NickName   string `bson:"nickname" json:"nickname"`
	SkillLevel int    `bson:"skilllevel" json:"skilllevel"`
}
