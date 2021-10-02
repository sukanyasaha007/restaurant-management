package models

import (
	"time"

	"go.mongodb-org/mongo-driver/bson/primitive"
)

type Table struct {
	ID               primitive.ObjectID `bson:"_id"`
	Number_of_guests *int               `json:"num_guests" validate:"required"`
	Table_num        *int               `json:"table_num" validate:"required"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
	Table_id         string             `json:"table_id" validate:"required"`
}
