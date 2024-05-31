package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"optician-rest-api/models/eyedata"
)

type Bills struct {
	Id                    primitive.ObjectID            `json:"id,omitempty" bson:"_id,omitempty"`
	Date                  string                        `json:"date,omitempty" bson:"date,omitempty"`
	CustomersBill         CustomersBill                 `json:"customersbill,omitempty" bson:"customersbill,omitempty"`
	SpectaclePrescription eyedata.SpectaclePrescription `json:"spectacle_prescription,omitempty" bson:"spectacle_prescription,omitempty"`
}
