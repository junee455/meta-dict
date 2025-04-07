package main

import (
	"meta-dict-back/dict_db"
	"meta-dict-back/handlers"
	"meta-dict-back/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func dbMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		initData := c.Request.Header.Get("InitData")
		dataValid, userDataQuery := utils.VerifyInitData(initData)

		if !dataValid {
			c.Abort()
			return
		}

		userData, err := utils.ParseUserDataFromQuery(userDataQuery)

		if err != nil {
			c.Abort()
			return
		}

		idStr := strconv.FormatInt(userData.ID, 10)

		conectionWithUser, err := dict_db.GetDBWithUser(idStr)

		if err != nil {
			c.Abort()
			return
		}

		c.Set("dbConnection", conectionWithUser)
	}
}

func main() {
	dict_db.InitDB()

	defer dict_db.Disconnect()

	router := gin.Default()

	router.GET("/healthcheck", handlers.Healthcheck)

	wordInfoGroup := router.Group("wordInfo")

	wordInfoGroup.Use(dbMiddleware())

	wordInfoGroup.GET("", handlers.GetWordsList)
	wordInfoGroup.GET(":word", handlers.GetWordInfo)
	wordInfoGroup.POST("", handlers.AddWordInfo)
	wordInfoGroup.PATCH("", handlers.UpdateWord)
	wordInfoGroup.DELETE(":word", handlers.DeleteWord)

	router.Run(":8080")
}
