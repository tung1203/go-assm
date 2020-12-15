package dal

import (
	"fmt"
	"quizz/model"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Register register
func Register(user model.User) (bool, string) {

	GetConnection()
	if !ValidateEmail(user.Email) {
		return false, "Invalid email"
	}
	exitsUser := db.First(&user)
	if exitsUser.RowsAffected > 0 {
		hashedPw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err == nil {
			user := model.User{
				Email:    user.Email,
				Password: string(hashedPw),
			}
			db.Create(&user)
			return true, "You have been successfully registered"

		} else {
			return false, "Error hashing"
		}

	} else {
		return false, "Email exists"
	}

}

//ValidateEmail regex email
func ValidateEmail(email string) bool {
	match, _ := regexp.MatchString(`(.+)@(.+)\.(.+)`, email)
	return match
}

var signingKey, verificationKey []byte

// Login login
func Login(user model.User) (bool, string, string, uint) {
	GetConnection()
	var foundUser model.User
	db.Where("email = ?", user.Email).First(&foundUser)
	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))

	if foundUser.ID > 0 && err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
			"iat":    time.Now().Unix(),
			"iss":    "admin",
			"alg":    "hs256",
			"userId": fmt.Sprint(foundUser.ID),
			"role":   "Member",
		})

		tokenString, err := token.SignedString([]byte("tungdz"))
		if err != nil {
			return false, "Error signing token", "", 0
		}
		foundUser.Password = ""
		return true, "Unauthorized", tokenString, foundUser.ID
	} else {
		return false, "Unauthorized", "", 0
	}
}

//GetUserByID get user by id
func GetUserByID(userID int64) (model.User, error) {
	GetConnection()
	var foundUser model.User
	db.Where("id = ?", userID).First(&foundUser)
	return foundUser, nil
}

//UpdateScore update score
func UpdateScore(userID int64, score int) bool {
	GetConnection()
	var foundUser model.User
	db.Where("id = ?", userID).First(&foundUser)

	newScore := foundUser.Score + score
	db.Table("users").Where("id = ?", userID).Update("score", newScore)
	return true
}

//Test get user by id
func Test() int {
	GetConnection()
	db.Create(&model.User{Email: "test@gmail.com", Password: "test"})
	return 1
}
