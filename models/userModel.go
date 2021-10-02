package models

import (
	"time"

	"go.mongodb-org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	User_id       string             `json:"user_id" validate:"required,min=6,max=20"`
	Password      *string            `json:"Password" validate:"required,min=6"`
	Avatar        *string            `json:"avatar"`
	Email         *string            `json:"email" validate:"required"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Phone         string             `json:"phone" validate:"required"`
	Order_id      *string            `json:"order_id"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}
