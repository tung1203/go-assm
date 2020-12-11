package main

import (
	"quizz/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/images", "./public/images")

	controllers.SetupQuestionRouter(r)
	controllers.SetupCategoryRouter(r)
	controllers.SetupTopicRouter(r)
	r.Run(":8080")
}
