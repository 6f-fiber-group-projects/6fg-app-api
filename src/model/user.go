package model

import (
	menty "6fg-app-api/entity/model_entity"
	// "fmt"
)

func GetAllUsers() ([]menty.User, error) {
	db := gormConnect()
	defer db.Close()

	users := []menty.User{}
	result := db.Find(&users)
	return users, result.Error
}

func GetUserById(userId int) (menty.User, error) {
	db := gormConnect()
	defer db.Close()

	user := menty.User{}
	result := db.Find(&user, "id=?", userId)
	return user, result.Error
}

func CreateUser(user *menty.User) (menty.User, error) {
	db := gormConnect()
	defer db.Close()

	result := db.Create(user)
	return *user, result.Error
}

func UpdateUser(u *menty.User) (menty.User, error) {
	db := gormConnect()
	defer db.Close()

	user := menty.User{}
	result := db.Model(&user).Where("id = ?", u.Id).Omit("id", "created_at").Updates(u)
	return user, result.Error
}

func DeleteUser(user *menty.User) (menty.User, error) {
	db := gormConnect()
	defer db.Close()

	result := db.Delete(user)
	return *user, result.Error
}
