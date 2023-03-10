package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID                     primitive.ObjectID `bson:"_id"`
	CreatedAt              string             `json:"created_at"`
	Name                   string             `json:"name"`
	Imgs                   []string           `json:"imgs"`
	Description            string             `json:"desc"`
	EstimatedPreparingTime *float32           `json:"est_prep_time"`
	Price                  *float32           `json:"price"`
}
