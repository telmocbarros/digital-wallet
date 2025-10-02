package database

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

var data []User = []User{
	{Id: "b18b851a-c8c4-4957-b68a-14362a1810c6", Email: "john@example.com", Password: "$2a$14$4Il8GoD6jpuFDi4ScOAqWuRZqK80cfZaUQ1TotEu2eDoIPFockbUC"}, // password123
	{Id: "b5ed9407-681b-4dbb-b2d3-997803e8bbfc", Email: "jane@example.com", Password: "$2a$14$Od/6Z6WvfnaRAFPlzsaEEuSgOfStbdAnBO20vpQYhjnK1TNzmJHmS"}, // securepass
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func VerifyUserCredentials(email, password string) (UserDTO, bool) {

	for _, user := range data {
		if user.Email == email && checkPasswordHash(password, user.Password) {
			return UserDTO{
				Email: user.Email,
				Id:    user.Id,
			}, true
		}
	}
	return UserDTO{}, false
}

func GetUserByEmail(email string) (UserDTO, bool) {
	for _, user := range data {
		if user.Email == email {
			return UserDTO{
				Email: user.Email,
				Id:    user.Id,
			}, true
		}
	}
	return UserDTO{}, false
}

func GetUserById(id string) (UserDTO, bool) {
	for _, user := range data {
		if user.Id == id {
			return UserDTO{
				Email: user.Email,
				Id:    user.Id,
			}, true
		}
	}
	return UserDTO{}, false
}

func GetUsers() []UserDTO {
	var users []UserDTO
	for _, user := range data {
		users = append(users, UserDTO{
			Email: user.Email,
			Id:    user.Id,
		})
	}

	return users
}

func SaveUser(password, email string) {
	newUserId := uuid.New().String()
	hashedPassword, _ := hashPassword(password)

	dbUser := User{
		Id:       newUserId,
		Email:    email,
		Password: hashedPassword,
	}
	data = append(data, dbUser)
}
