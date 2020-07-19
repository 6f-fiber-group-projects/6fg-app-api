package businessLogic

import (
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	repo "github.com/6f-fiber-group-projects/6fg-app-api/repository"
	"fmt"
)

func UpdateEquipmentStatus(r *reqenty.EquipmentHistoryRequest) error {
	// validate user
	_, err := repo.GetUserById(r.UserId)
	if err != nil {
		return fmt.Errorf("No user was found (id: %d)", r.UserId)
	}

	// validate equip
	equip, err := repo.GetEquipmentById(r.EquipId)
	if err != nil {
		return fmt.Errorf("No equipment was found (id: %d)", r.EquipId)
	}

	// validate reservation
	if r.ReservationId != nil {
		_, err = repo.GetEquipmentReservationById(*r.ReservationId)
		if err != nil {
			return fmt.Errorf("No reservation was found (id: %d)", *r.ReservationId)
		}
	}

	// check equip_sataus and update current history to end
	equipUpdate := equipModelToReq(&equip)
	if r.EquipStatus == 0 || *equip.Status == 1 {
		equipUpdate.Status = 0
		_, err = repo.UpdateEquipment(&equipUpdate)
		if err != nil {
			return fmt.Errorf("%s", err)
		}

		currentEquipHistroy, err := repo.GetLatestEquipmentHistoryByEquipId(equip.Id)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		currentEquipHistroyUpdate := equipHistroyModelToReq(&currentEquipHistroy)

		_, err = repo.UpdateEquipmentHistory(&currentEquipHistroyUpdate)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
	}

	// create new history to start
	if r.EquipStatus == 1 {
		equipUpdate.Status = 1
		_, err = repo.UpdateEquipment(&equipUpdate)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		_, err = repo.CreateEquipmentHistory(r)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
	}

	return nil
}

func equipModelToReq(e *menty.Equipment) reqenty.EquipmentUpdateRequest {
	return reqenty.EquipmentUpdateRequest{
		Id:     e.Id,
		Status: *e.Status,
		EquipmentRequest: &reqenty.EquipmentRequest{
			Name: e.Name,
		},
	}
}

func equipHistroyModelToReq(e *menty.EquipmentHistory) reqenty.EquipmentHistoryUpdateRequest {
	return reqenty.EquipmentHistoryUpdateRequest{
		Id: e.Id,
		EquipmentHistoryRequest: &reqenty.EquipmentHistoryRequest{
			EquipId:       e.EquipId,
			UserId:        e.UserId,
			ReservationId: e.ReservationId,
		},
	}
}

// add scheuler for the case user forget to stop using equip
