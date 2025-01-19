package main

import (
	"meta-dict-back/dict_db"
	"meta-dict-back/handlers"
	// "meta-dict-back/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	dict_db.InitDB()

	defer dict_db.Disconnect()

	// cors middleware for frontend
	// migrations.MigrateWords()

	router := gin.Default()

	router.GET("/wordInfo", handlers.GetWordsList)
	router.GET("/wordInfo/:word", handlers.GetWordInfo)
	router.POST("/wordInfo", handlers.AddWordInfo)
	router.PATCH("/wordInfo", handlers.UpdateWord)
	router.DELETE("/wordInfo/:wordId", handlers.DeleteWord)

	router.Run(":8080")
}
