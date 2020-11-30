package businessLogic

import (
	"fmt"
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	repo "github.com/6f-fiber-group-projects/6fg-app-api/repository"
	"github.com/google/uuid"
)

func GetEquipmentReservation() ([]menty.EquipmentReservation, error) {
	rsvns, err := repo.GetEquipmentReservation()
	if err != nil {
		return nil, err
	}
	return rsvns, nil
}

func ReserveEquip(r *[]reqenty.EquipmentReservationRequest) error {
	// validate user
	_, err := repo.GetUserById((*r)[0].UserId)
	if err != nil {
		return fmt.Errorf("No user was found (id: %d)", (*r)[0].UserId)
	}

	// validate equip
	for _, rsvn := range *r {
		_, err = repo.GetEquipmentById(rsvn.EquipId)
		if err != nil {
			return fmt.Errorf("No equipment was found (id: %d)", rsvn.EquipId)
		}
	}

	uu, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	for _, rsvn := range *r {
		rsvn.GroupId = uu.String()
		_, err = repo.CreateEquipmentReservation(&rsvn)
		if err != nil {
			return err
		}
	}
	return nil
}
