package repository

import (
	menty "6fg-app-api/entity/model_entity"
	reqenty "6fg-app-api/entity/request_entity"
	model "6fg-app-api/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetAllUsers() ([]menty.User, error) {
	users, err := model.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return users, nil
}

func GetUserById(userId int) (menty.User, error) {
	user, err := model.GetUserById(userId)
	if err != nil {
		return menty.User{}, fmt.Errorf("%s", err)
	}
	return user, nil
}

func CreateUser(u *reqenty.UserRequest) (menty.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	user := menty.User{
		Authority_id: u.Authority_id,
		Google_id:    u.Google_id,
		Name:         u.Name,
		Email:        u.Email,
		Password:     hash,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
	}

	user, err := model.CreateUser(&user)
	if err != nil {
		return menty.User{}, fmt.Errorf("%s", err)
	}

	// err = bcrypt.CompareHashAndPassword(hash, []byte(u.Password))
	// if err != nil {
	// 	fmt.Println("Failure")
	// } else {
	// 	fmt.Println("Success")
	// }

	return user, nil
}

func UpdateUser(u *reqenty.UserUpdateRequest) (menty.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	user := menty.User{
		Id:           u.Id,
		Authority_id: u.Authority_id,
		Google_id:    u.Google_id,
		Name:         u.Name,
		Email:        u.Email,
		Password:     hash,
		Updated_at:   time.Now(),
	}

	user, err := model.UpdateUser(&user)
	if err != nil {
		return menty.User{}, fmt.Errorf("%s", err)
	}

	return user, nil
}

func DeleteUser(userId int) (menty.User, error) {
	user, err := model.GetUserById(userId)
	if err != nil {
		return menty.User{}, fmt.Errorf("%s", err)
	}

	_, err = model.DeleteUser(&user)
	if err != nil {
		return menty.User{}, fmt.Errorf("%s", err)
	}

	return user, nil
}
