package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mflix struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Year        int                `json:"year" bson:"year"`
	Director    string             `json:"director" bson:"director"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

type Comment struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Email   string             `json:"email" bson:"email"`
	MovieID primitive.ObjectID `json:"movie_id" bson:"movie_id"`
	Text    string             `json:"text" bson:"text"`
	Date    primitive.DateTime `json:"date" bson:"date"`
}
