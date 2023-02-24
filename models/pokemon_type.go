package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PokemonType struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Name       string             `json:"name,omitempty" validate:"required"`
	Stengths   []string           `json:"stengths,omitempty"`
	Weaknesses []string           `bson:"weaknesses,omitempty"`
	ImmuneTo   []string           `bson:"immuneTo,omitempty"`
	CannotHit  []string           `bson:",cannotHit,omitempty"`
}
