package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title string             `bson:"title" json:"title" binding:"required"`
	Done  bool               `bson:"done" json:"done"`
}
