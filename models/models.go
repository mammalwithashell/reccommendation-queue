package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Recomendation to be stored in the database
type Recomendation struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Medium string             `bson:"medium,omitempty"`
}

// User struct to hold user information
type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	UserID         string             `bson:"user_id,omitempty"`
	Name           string             `bson:"name,omitempty"`
	Email          string             `bson:"email,omitempty"`
	HashedPassword string             `bson:"hashedPassword,omitempty"`
	Authenticated  bool               `bson:"auth,omitempty"`
}
