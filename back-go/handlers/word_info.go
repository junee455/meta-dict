package handlers

import (
	"context"
	"fmt"
	"net/http"

	"meta-dict-back/dict_db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetWordsList(c *gin.Context) {
	db := dict_db.GetDB()

	collections, _ := db.DB.ListCollectionNames(context.Background(), bson.D{})

	fmt.Println("GET WORDS")
	fmt.Printf("Collections: %v \n", len(collections))

	words, err := db.GetWordsList()

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, words)
}

func GetWordInfo(c *gin.Context) {
	db := dict_db.GetDB()

	req_word := c.Params.ByName("word")

	word, err := db.FindWord(req_word)

	if err != nil {
		c.Status(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, word)
}

func UpdateWord(c *gin.Context) {
	db := dict_db.GetDB()

	var word_data dict_db.WordSchema

	// get word json
	if err := c.ShouldBindJSON(&word_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	result, err := db.UpdateWord(word_data.ID, word_data)

	fmt.Printf("%v %v\n", result, err)

	if err != nil || result.ModifiedCount != 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update word"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func DeleteWord(c *gin.Context) {
	db := dict_db.GetDB()
	word_id_param := c.Params.ByName("wordId")

	word_id, err := primitive.ObjectIDFromHex(word_id_param)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid id"})
		return
	}

	db_res, err := db.DeleteWord(word_id)

	if err != nil || db_res.DeletedCount != 1 {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func AddWordInfo(c *gin.Context) {
	db := dict_db.GetDB()

	var word_data dict_db.WordSchema

	// get word json
	if err := c.ShouldBindJSON(&word_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}

	newID, err := db.AddNewWord(word_data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert new word"})
	}

	c.JSON(http.StatusOK, newID)

	// db.AddNewWord(dict_db.WordSchema{
	// Word: ,
	// })
}
