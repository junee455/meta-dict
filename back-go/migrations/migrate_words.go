package migrations

import (
	"context"
	"fmt"
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
	// user that owns the word
	Owner primitive.ObjectID `bson:"owner,omitempty" json:"owner,omitempty"`
}

type UserSchema struct {
	ID    primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"` // MongoDB automatically sets this field if omitted
	TgID  string               `bson:"tgID" json:"tgID"`
	Words []dict_db.WordSchema `bson:"words" json:"words"`
}

func MigrateWords(tgUserId string) {
	db := dict_db.GetDB()

	filter := bson.M{}

	ctx := context.Background()

	wordsCollection := db.DB.Collection("words")

	usersCollection := db.DB.Collection("users")

	var words []dict_db.WordSchema

	cursor, err := wordsCollection.Find(ctx, filter)

	if err != nil {
		return
	}

	fmt.Println("created new user doc")

	err = cursor.All(context.TODO(), &words)

	newUserDoc := bson.M{"tgID": tgUserId, "words": words}

	if err != nil {
		return
	}

	fmt.Println("created new user doc")

	_, err = usersCollection.InsertOne(context.TODO(), newUserDoc)

	if err != nil {
		return
	}
}
