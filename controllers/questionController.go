package controllers

import (
	"quizz/dal"
	"quizz/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SetupQuestionRouter for item router
func SetupQuestionRouter(r *gin.Engine) {
	//Create
	r.POST("/questions", func(c *gin.Context) {
		var question model.Question
		if err := c.ShouldBindJSON(&question); err == nil {
			rowsAffected, lastInsertedID, err := dal.InsertQuestion(question)
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "Insert Item error",
				})
			} else {
				if rowsAffected > 0 {
					c.JSON(200, gin.H{
						"messages": "Insert Item complete",
						"itemId":   lastInsertedID,
					})
				}
			}
		}
	})

	// //Read
	r.GET("/questions/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		question, err := dal.GetQuestion(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Question not found",
			})
		} else {
			c.JSON(200, question)
		}
	})

	r.POST("/questions/checkAnswers", func(c *gin.Context) {
		var answer []model.Answer
		if err := c.ShouldBindJSON(&answer); err == nil {
			numberAnswerCorrect, err := dal.CheckAnswers(answer)
			if err != nil {
				c.JSON(200, gin.H{
					"messages": "Something went wrong",
				})
			} else {
				c.JSON(200, numberAnswerCorrect)

			}
		}
	})
	//Update
	// r.PUT("/item", func(c *gin.Context) {
	// 	var item model.Item
	// 	if err := c.ShouldBindJSON(&item); err == nil {
	// 		rowsAffected, err := dal.UpdateItem(item)
	// 		if err != nil {
	// 			c.JSON(500, gin.H{
	// 				"messages": "update Item error",
	// 			})
	// 		} else {
	// 			if rowsAffected > 0 {
	// 				c.JSON(200, gin.H{
	// 					"messages": "update Item complete",
	// 				})
	// 			}
	// 		}
	// 	}
	// })

	//Delete
	// r.DELETE("/item/:id", func(c *gin.Context) {
	// 	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	// 	rowsDeletedAffected, err := dal.DeleteItem(id)
	// 	if err != nil {
	// 		c.JSON(500, gin.H{
	// 			"messages": "delete error.",
	// 		})
	// 	} else {
	// 		if rowsDeletedAffected > 0 {
	// 			c.JSON(200, gin.H{
	// 				"messages": "delete completed.",
	// 			})
	// 		}
	// 	}
	// })
}
