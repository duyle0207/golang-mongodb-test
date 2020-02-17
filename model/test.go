package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type test struct {
	id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	content string `json:"content,omitempty" bson:"content,omitempty"`
}

