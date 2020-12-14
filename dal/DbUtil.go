package dal

import (
	"fmt"
	"os"
	"quizz/model"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//InitializeMySQL to OrderDB
func InitializeMySQL() {
	dsn1 := os.Getenv("DATABASE_URL")

	// dBConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("Error opening database: %q", err)
	// }
	// dsn := "root:tung1203@tcp(127.0.0.1:3306)/quizz?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "mysql://bc3e4c2f9def1b:c45e50e7@us-cdbr-east-02.cleardb.com/heroku_5598edf86a828a0?reconnect=true"
	println(dsn)
	println(dsn1)
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
