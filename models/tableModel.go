package models

import (
	"time"

	"go.mongodb-org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_date time.Time          `json:"created_at" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Order_id   string             `json:"order_id" validate:"required"`
	Table_id   string             `json:"table_id" validate:"required"`
}
