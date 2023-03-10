package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID           primitive.ObjectID `bson:"_id"`
	CreationDate string             `json:"creation_date"`
	ConfirmDate  string             `json:"confirm_date`
	Token        string             `json:"token"`
	SenderName   string             `json:"sender_name"`
	Amount       float32            `json:"amount"`
}
