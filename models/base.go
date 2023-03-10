package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
	DeletedAt string             `json:"deleted_at"`
}
