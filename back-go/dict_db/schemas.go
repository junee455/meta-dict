package dict_db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WordSchema struct {
	Word         string               `bson:"word" json:"word"`
	Translations []string             `bson:"translations" json:"translations"`
	Description  string               `bson:"description" json:"description"`
	Metadata     string               `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Similar      []primitive.ObjectID `bson:"similar,omitempty" json:"similar,omitempty"`
}

type UserSchema struct {
	ID    string       `bson:"_id,omitempty" json:"id,omitempty"`
	TgID  string       `bson:"tgID" json:"tgID"`
	Words []WordSchema `bson:"words" json:"words"`
}
