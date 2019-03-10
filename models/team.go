package models

// Team represents a document in the teams collection
type Team struct {
	ID     string   `json:"_id"`
	Name   string   `json:"name"`
	Member []string `json:"member"`
}
