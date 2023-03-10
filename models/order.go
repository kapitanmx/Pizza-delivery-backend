package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `bson:"_id"`
	Purchaser     *User              `json:"purchaser"`
	TimeTotal     *float32           `json:"time_total"`
	TimeRemaining *float32           `json:"time_remaining"`
	TimePassed    *float32           `json:"time_passed"`
	Products      []Product          `json:"products"`
	OrderPrice    *float32           `json:"order_price"`
	DeliveryPrice *float32           `json:"delivery_price"`
	TotalPrice    *float32           `json:"total_price"`
}
