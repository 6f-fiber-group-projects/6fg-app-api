package businessLogic

import (
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	repo "github.com/6f-fiber-group-projects/6fg-app-api/repository"
	"fmt"
)

func ReserveEquip(r *reqenty.EquipmentReservationRequest) error {
	// validate user
	_, err := repo.GetUserById(r.UserId)
	if err != nil {
		return fmt.Errorf("No user was found (id: %d)", r.UserId)
	}

	// validate equip
	_, err = repo.GetEquipmentById(r.EquipId)
	if err != nil {
		return fmt.Errorf("No equipment was found (id: %d)", r.EquipId)
	}

	_, err = repo.CreateEquipmentReservation(r)
	if err != nil {
		return err
	}

	return nil
}
