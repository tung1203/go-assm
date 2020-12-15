package main

import (
	"os"
	"quizz/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/images", "./public/images")

	controllers.SetupUserRouter(r)
	controllers.SetupQuestionRouter(r)
	controllers.SetupCategoryRouter(r)
	controllers.SetupTopicRouter(r)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	r.Run(":" + port)
}
