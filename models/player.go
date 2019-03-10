package models

// Player represents a Person who want's to play
type Player struct {
	ID         string `json:"_id"`
	LastName   string `json:"lastname"`
	FirstName  string `json:"firstname"`
	NickName   string `json:"nickname"`
	SkillLevel int    `json:"skilllevel"`
}
