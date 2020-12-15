package controllers

import (
	"fmt"
	"quizz/dal"
	"quizz/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SetupUserRouter for item router
func SetupUserRouter(r *gin.Engine) {
	// //Read
	r.POST("/login", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err == nil {
			success, err, token := dal.Login(user)
			fmt.Println(user.Email)
			if success == true {
				c.JSON(200, gin.H{
					"messages": err,
					"token":    token,
				})
			} else {
				c.JSON(500, gin.H{
					"messages": err,
				})
			}
		} else {
			c.JSON(500, gin.H{
				"messages": "Nothing",
			})
		}
	})

	r.POST("/register", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err == nil {
			success, erra := dal.Register(user)

			if success == true {
				c.JSON(200, gin.H{
					"messages": erra,
				})
			} else {
				c.JSON(500, gin.H{
					"messages": erra,
				})

			}
		}
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		user, err := dal.GetUserByID(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "User not found",
			})
		} else {
			c.JSON(200, user)
		}
	})

	type PostScoreModel struct {
		ID    int64 `json:"id"`
		Score int   `json:"score"`
	}

	r.POST("/users/score", func(c *gin.Context) {
		var postScore PostScoreModel
		if err := c.ShouldBindJSON(&postScore); err == nil {
			success := dal.UpdateScore(postScore.ID, postScore.Score)
			if success == true {
				c.JSON(200, gin.H{
					"messages": success,
				})
			} else {
				c.JSON(500, err)

			}
		}

	})
	r.GET("/usertest", func(c *gin.Context) {
		dal.Test()
		c.JSON(200, gin.H{
			"messages": "ok",
		})
	})

}
