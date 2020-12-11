package dal

import (
	"quizz/model"
)

// GetTopics get all topic
func GetTopics() ([]model.Topic, error) {
	GetConnection()

	var topics []model.Topic

	db.Limit(10).Find(&topics)

	return topics, nil
}

//GetTopicByCategoryID get all topics by category id
func GetTopicByCategoryID(categoryID int64) ([]model.Topic, error) {
	GetConnection()

	var topics []model.Topic

	db.Where("category_id = ?", categoryID).Find(&topics)

	return topics, nil
}
