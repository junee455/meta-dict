package handlers

import (
	"context"
	"net/http"

	"meta-dict-back/dict_db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func updateUserWordList(connection *dict_db.ConnectionWithUser, words []dict_db.WordSchema) (*mongo.UpdateResult, error) {
	update := bson.M{"$set": bson.M{
		"words": words,
	},
	}

	idHex, _ := primitive.ObjectIDFromHex(connection.User.ID)

	res, err := connection.UsersCollection.UpdateByID(context.TODO(), idHex, update)

	return res, err
}

func GetWordsList(c *gin.Context) {
	connectionAny, exists := c.Get("dbConnection")

	if !exists {
		c.Status(http.StatusInternalServerError)

		return
	}

	connection, ok := connectionAny.(*dict_db.ConnectionWithUser)

	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, connection.User.Words)
}

func GetWordInfo(c *gin.Context) {
	req_word := c.Params.ByName("word")

	connectionAny, exists := c.Get("dbConnection")

	if !exists {
		c.Status(http.StatusInternalServerError)

		return
	}

	connection, ok := connectionAny.(*dict_db.ConnectionWithUser)

	if !ok {
		c.Status(http.StatusInternalServerError)

		return
	}

	for _, word := range connection.User.Words {
		if word.Word == req_word {
			c.JSON(http.StatusOK, word)
			return
		}
	}

	c.Status(http.StatusNotFound)
}

func UpdateWord(c *gin.Context) {
	// validate connection >>

	connectionAny, exists := c.Get("dbConnection")

	if !exists {
		c.Status(http.StatusInternalServerError)

		return
	}

	connection, ok := connectionAny.(*dict_db.ConnectionWithUser)

	if !ok {
		c.Status(http.StatusInternalServerError)

		return
	}

	// validate connection <<

	var word_data dict_db.WordSchema

	// get word json
	if err := c.ShouldBindJSON(&word_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	req_word := word_data.Word

	// find word
	var wordIndex int = -1

	for i, w := range connection.User.Words {
		if w.Word == req_word {
			wordIndex = i
			break
		}
	}

	if wordIndex == -1 {
		c.Status(http.StatusInternalServerError)
		return
	}

	connection.User.Words[wordIndex] = word_data

	res, err := updateUserWordList(connection, connection.User.Words)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if res.MatchedCount == 0 {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func DeleteWord(c *gin.Context) {
	connectionAny, exists := c.Get("dbConnection")

	if !exists {
		c.Status(http.StatusInternalServerError)

		return
	}

	connection, ok := connectionAny.(*dict_db.ConnectionWithUser)

	if !ok {
		c.Status(http.StatusInternalServerError)

		return
	}

	req_word := c.Params.ByName("word")

	var wordIndex int = -1

	for i, w := range connection.User.Words {
		if w.Word == req_word {
			wordIndex = i
			break
		}
	}

	if wordIndex == -1 {
		c.Status(http.StatusInternalServerError)
		return
	}

	connection.User.Words = append(connection.User.Words[:wordIndex], connection.User.Words[wordIndex+1:]...)

	res, err := updateUserWordList(connection, connection.User.Words)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if res.MatchedCount == 0 {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func AddWordInfo(c *gin.Context) {
	connectionAny, exists := c.Get("dbConnection")

	if !exists {
		c.Status(http.StatusInternalServerError)

		return
	}

	connection, ok := connectionAny.(*dict_db.ConnectionWithUser)

	if !ok {
		c.Status(http.StatusInternalServerError)

		return
	}

	// get word json
	var word_data dict_db.WordSchema

	if err := c.ShouldBindJSON(&word_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}

	if len(connection.User.Words) < 1 {
		connection.User.Words = []dict_db.WordSchema{}
	}

	connection.User.Words = append(connection.User.Words, word_data)

	res, err := updateUserWordList(connection, connection.User.Words)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if res.MatchedCount == 0 {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
