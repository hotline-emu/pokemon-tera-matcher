package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PokemonType struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"text"`
	Stengths   []string           `bson:"array"`
	Weaknesses []string           `bson:"array"`
	ImmuneTo   []string           `bson:"array"`
	CannotHit  []string           `bson:"array"`
}
