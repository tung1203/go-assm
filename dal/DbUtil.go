package dal

import (
	"fmt"
	"quizz/model"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//InitializeMySQL to OrderDB
func InitializeMySQL() {
	// dsn := os.Getenv("DATABASE_URL")
	// dBConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("Error opening database: %q", err)
	// }
	dsn := "root:tung1203@tcp(127.0.0.1:3306)/quizz?charset=utf8mb4&parseTime=True&loc=Local"
	dBConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connect Successfull.")

	db = dBConnection
	db.AutoMigrate(&model.Category{}, &model.Topic{}, &model.Question{}, &model.Answer{})
}

//GetConnection is get MySQL Connection
func GetConnection() *gorm.DB {
	if db == nil {
		InitializeMySQL()
	}
	return db
}
