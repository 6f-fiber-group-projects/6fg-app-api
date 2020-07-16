package businessLogic

import (
	menty "6fg-app-api/entity/model_entity"
	reqenty "6fg-app-api/entity/request_entity"
	repo "6fg-app-api/repository"
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

	// check equip_sataus and update current history to end
	equipUpdate := equipHistroyModelToReq(&equip)
	if r.EquipStatus == 0 || *equip.Status == 1 {
		equipUpdate.Status = 0
		_, err = repo.UpdateEquipment(&equipUpdate)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		_, err = repo.UpdateEquipmentHistory(r)
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

func equipHistroyModelToReq(e *menty.Equipment) reqenty.EquipmentUpdateRequest {
	return reqenty.EquipmentUpdateRequest{
		Id:     e.Id,
		Status: *e.Status,
		EquipmentRequest: &reqenty.EquipmentRequest{
			Name: e.Name,
		},
	}
}

// add scheuler for the case user forget to stop using equip
