package models

import "gopkg.in/mgo.v2/bson"

// Represents a user, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
	Age         int           `bson:"age" json:"age"`
	Gender      string        `bson:"gender" json:"gender"`
}
