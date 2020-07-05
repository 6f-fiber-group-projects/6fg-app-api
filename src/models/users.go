package models

import (
	// "golang.org/x/crypto/bcrypt"
	menty "6fg-app-api/entities/model_entities"
	"fmt"
)

func GetAllUsers() {

}

func GetUserById() {

}

func CreateUser(u *menty.User) {
	db := gormConnect()
	defer db.Close()

	err := db.Create(u)
	fmt.Println(err)
	// hash, err := bcrypt.GenerateFromPassword([]byte("plain text"), 8)
}
