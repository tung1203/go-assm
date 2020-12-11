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

	// item := model.Item{0, "Item 100", 15.5, 12, 1, ""}
	// rowsAffected, lastInsertedID, err := dal.InsertItem(item)

	// if err != nil {
	// 	fmt.Println("Insert error.")
	// } else {
	// 	if rowsAffected > 0 {
	// 		fmt.Println("Insert completed.")
	// 		item.ItemId = lastInsertedID
	// 		fmt.Printf("new item id is %d\n", item.ItemId)
	// 	}
	// }

	// //update item
	// item.ItemName = "Item 10001"
	// rowsUpdatedAffected, err := dal.UpdateItem(item)
	// if err != nil {
	// 	fmt.Println("Update error.")
	// } else {
	// 	if rowsUpdatedAffected > 0 {
	// 		fmt.Println("update completed.")
	// 	}
	// }

	// //delete item
	// rowsDeletedAffected, err := dal.DeleteItem(lastInsertedID)
	// if err != nil {
	// 	fmt.Println("delete error.")
	// } else {
	// 	if rowsDeletedAffected > 0 {
	// 		fmt.Println("delete completed.")
	// 	}
	// }
}
