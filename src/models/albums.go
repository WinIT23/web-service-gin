package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	ID     primitive.ObjectID `json:"_id,omitempty"`
	Title  string             `json:"title" validate:"required"`
	Artist string             `json:"artist" validate:"required"`
	Price  float64            `json:"price" validate:"required"`
}
