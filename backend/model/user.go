package model

import "time"

const (
	usersTableName = "users"
)

// User represents the model of the users database table.
type User struct {
	ID        int       `json:"id" sql:"primary_key"`
	Name      string    `json:"name" sql:"type:varchar(32);unique;not null"`
	Password  []byte    `json:"-" sql:"type:bytea;not null"`
	Email     string    `json:"email" sql:"type:varchar(256);unique;not null"`
	CreatedAt time.Time `json:"created_at" sql:"not null;default:now()"`
	UpdatedAt time.Time `json:"updated_at" sql:"not null;default:now()"`
}

// TableName returns the name of the table.
func (User) TableName() string {
	return usersTableName
}
