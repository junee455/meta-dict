package migrations

import (
	"context"
	"meta-dict-back/dict_db"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewWordSchema struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"` // MongoDB automatically sets this field if omitted
	Word         string               `bson:"word" json:"word"`
	Translations []string             `bson:"translations" json:"translations"`
	Description  string               `bson:"description" json:"description"`
	Metadata     string               `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Similar      []primitive.ObjectID `bson:"similar,omitempty" json:"similar,omitempty"`
}

func MigrateWords() {
	db := dict_db.GetDB()

	filter := bson.M{}

	ctx := context.Background()

	collection := db.DB.Collection("words")

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return
	}

	for cursor.Next(ctx) {
		var oldWord dict_db.WordSchema
		if err := cursor.Decode(&oldWord); err != nil {
			return
		}

		// var newWord = NewWordSchema{
		// 	ID:           oldWord.ID,
		// 	Word:         oldWord.Word,
		// 	Translations: strings.Split(oldWord.Translation, ","),
		// 	Description:  oldWord.Description,
		// 	Metadata:     oldWord.Metadata,
		// 	Similar:      oldWord.Similar,
		// }

		_, err := collection.UpdateByID(ctx, oldWord.ID, bson.M{
			"$unset": bson.M{
				"translation": "",
			},
		})

		if err != nil {
			return
		}
	}
}
