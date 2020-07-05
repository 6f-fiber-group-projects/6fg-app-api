package repositories

import (
	menty "6fg-app-api/entities/model_entities"
	renty "6fg-app-api/entities/request_entities"
	model "6fg-app-api/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func CreateUser(u *renty.UserRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return fmt.Errorf("encypt error")
	}

	user := menty.User{
		Authority_id: u.Authority_id,
		Google_id:    u.Google_id,
		Name:         u.Name,
		Email:        u.Email,
		Password:     hash,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
	}

	model.CreateUser(&user)

	return nil
}
