package controllers

import (
	"quizz/dal"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SetupCategoryRouter for item router
func SetupCategoryRouter(r *gin.Engine) {
	// //Read
	r.GET("/categories", func(c *gin.Context) {
		categories, err := dal.GetCategories()
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Category not found",
			})
		} else {
			c.JSON(200, categories)
		}
	})
	r.GET("/categories/:id/topics", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		topics, err := dal.GetTopicByCategoryID(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Topic not found",
			})
		} else {
			c.JSON(200, topics)
		}
	})

}
