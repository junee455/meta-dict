package dict_db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WordSchema struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"` // MongoDB automatically sets this field if omitted
	Word         string               `bson:"word" json:"word"`
	Translations []string             `bson:"translations" json:"translations"`
	Description  string               `bson:"description" json:"description"`
	Metadata     string               `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Similar      []primitive.ObjectID `bson:"similar,omitempty" json:"similar,omitempty"`
}
