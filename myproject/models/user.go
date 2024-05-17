package models

// User represents a registered user
type User struct {
	Nickname          string `json:"nickname" bson:"nickname"`
	Email             string `json:"email" bson:"email"`
	Password          string `json:"password" bson:"password"`
	Confirmed         bool   `json:"confirmed" bson:"confirmed"`
	ConfirmationToken string `json:"-" bson:"confirmationToken"`
}
