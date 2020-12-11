package dal

import (
	"fmt"
	"quizz/model"
)

//InsertQuestion to OrderDB rowsAffected, lastInsertID, err
func InsertQuestion(question model.Question) (int64, uint, error) {
	GetConnection()
	fmt.Println(question.Content)

	fmt.Println(question.Answers)
	newQuestion := model.Question{Content: question.Content, Answers: question.Answers}

	result := db.Create(&newQuestion) // pass pointer of data to Create

	return result.RowsAffected, newQuestion.ID, result.Error
}

//GetQuestion from itemId
func GetQuestion(questionID int64) (model.Question, error) {
	GetConnection()

	var question model.Question

	db.Preload("Answers").First(&question, questionID)
	return question, nil

}

// GetQuestionByTocpicID get questions by topic id
func GetQuestionByTocpicID(topicID int64) (model.Topic, error) {
	GetConnection()
	var topic model.Topic
	db.Preload("Questions").First(&topic, topicID)
	return topic, nil
}

// CheckAnswers check answers by question Id
func CheckAnswers(answers []model.Answer) (int, error) {
	GetConnection()
	// answers[0].ID
	var numberAnswerCorrected int = 0
	for i, answer := range answers {
		fmt.Printf("%d. Id: %d questionId: %d", i, answer.ID, answer.QuestionID)
		var answerDb model.Answer
		db.Where("id = ? AND question_id >= ?", answer.ID, answer.QuestionID).Find(&answerDb)
		println(answerDb.IsCorrect)
		if answerDb.IsCorrect {
			numberAnswerCorrected++
		}
	}
	fmt.Println(numberAnswerCorrected)
	return numberAnswerCorrected, nil
}
