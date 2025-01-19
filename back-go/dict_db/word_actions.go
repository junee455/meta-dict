package dict_db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (dc *DictDBConnection) AddNewWord(word WordSchema) (primitive.ObjectID, error) {
	words_collectoin := dc.DB.Collection("words")

	word.ID = primitive.NewObjectID()

	insertRes, err := words_collectoin.InsertOne(context.Background(), word)

	return insertRes.InsertedID.(primitive.ObjectID), err
}

func (dc *DictDBConnection) UpdateWord(id primitive.ObjectID, word WordSchema) (*mongo.UpdateResult, error) {
	update := bson.M{"$set": word}

	return dc.DB.Collection("words").UpdateByID(context.Background(), id, update)
}

func (dc *DictDBConnection) DeleteWord(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	return dc.DB.Collection("words").DeleteOne(context.Background(), filter)
}

func (dc *DictDBConnection) FindWord(word string) (WordSchema, error) {
	var result WordSchema

	err := dc.DB.Collection("words").FindOne(context.TODO(), bson.D{{"word", word}}).Decode(&result)

	return result, err
}

func (dc *DictDBConnection) GetAllWords() {}

func (dc *DictDBConnection) GetWordsList() ([]WordSchema, error) {
	cursor, err := dc.DB.Collection("words").Find(context.Background(), bson.M{})

	var words []WordSchema

	for cursor.Next(context.Background()) {
		var word WordSchema
		cursor.Decode(&word)
		words = append(words, word)
	}

	return words, err
}
