package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         *string            `json:"name"`
	LastName     *string            `json:"last_name"`
	Email        *string            `json:"email"`
	Password     *string            `json:"password"`
	PhoneNumber  *string            `json:"phone"`
	Street       *string            `json:"street"`
	HouseNumber  *string            `json:"house_number"`
	PostalCode   *string            `json:"postal_code"`
	City         *string            `json:"city"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserID       string             `json:"user_id"`
}
