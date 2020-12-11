package controllers

import (
	"quizz/dal"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SetupTopicRouter for item router
func SetupTopicRouter(r *gin.Engine) {
	// //Read
	r.GET("/topics", func(c *gin.Context) {
		topics, err := dal.GetTopics()
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Topics not found",
			})
		} else {
			c.JSON(200, topics)
		}
	})

	r.GET("/topics/:id/questions/", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		questions, err := dal.GetQuestionByTocpicID(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Questions not found",
			})
		} else {
			c.JSON(200, questions)
		}
	})

}
