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
	dsn := os.Getenv("JAWSDB_MARIA_URL")

	// dBConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("Error opening database: %q", err)
	// }
	// dsn := "root:tung1203@tcp(127.0.0.1:3306)/quizz?charset=utf8mb4&parseTime=True&loc=Local"
	dBConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connect Successfull.")

	db = dBConnection
	db.AutoMigrate(&model.Category{}, &model.Topic{}, &model.Question{}, &model.Answer{})
}

// mysql -h klbcedmmqp7w17ik.cbetxkdyhwsb.us-east-1.rds.amazonaws.com -u ow6qah1qqqgfqs80 -p l4xcudjgm2tc5qlq < db.sql
//GetConnection is get MySQL Connection
func GetConnection() *gorm.DB {
	if db == nil {
		InitializeMySQL()
	}
	return db
}
