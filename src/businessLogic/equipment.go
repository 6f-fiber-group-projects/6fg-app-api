package businessLogic

import (
	"fmt"
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	"github.com/6f-fiber-group-projects/6fg-app-api/lib"
	repo "github.com/6f-fiber-group-projects/6fg-app-api/repository"
	"github.com/gin-gonic/gin"
	"os"
)

func CreateEquipment(c *gin.Context, e *reqenty.EquipmentRequest) (*menty.Equipment, error) {
	if !IsAdmin(c) {
		return nil, fmt.Errorf("unauthorized")
	}

	equip, err := repo.CreateEquipment(e)
	if err != nil {
		return nil, err
	}

	_, err = repo.CreateEquipmentStatus(equip.Id)
	if err != nil {
		return nil, err
	}

	return &equip, nil
}

func UpdateEquipment(c *gin.Context, e *reqenty.EquipmentUpdateRequest) (*menty.Equipment, error) {
	if !IsAdmin(c) {
		return nil, fmt.Errorf("unauthorized")
	}

	equip, err := repo.UpdateEquipment(e)
	if err != nil {
		return nil, err
	}

	return &equip, nil
}

func DeleteEquipment(c *gin.Context, equipId int) (*menty.Equipment, error) {
	if !IsAdmin(c) {
		return nil, fmt.Errorf("unauthorized")
	}

	equip, err := repo.DeleteEquipment(equipId)
	if err != nil {
		return nil, err
	}

	return &equip, nil
}

func GenerateEquipmentQR(equipId int) ([]byte, error) {
	qr, err := lib.GenerateQR(fmt.Sprintf("%s/equipment/%d", os.Getenv("DOMEIN"), equipId))
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return qr, nil
}

func UpdateEquipmentStatus(c *gin.Context, r *reqenty.EquipmentHistoryRequest) error {
	return nil
}

// func UpdateEquipmentStatus(c *gin.Context, r *reqenty.EquipmentHistoryRequest) error {
// 	if !IsAdmin(c) && !IsSameUser(c, r.UserId) {
// 		return fmt.Errorf("unauthorized")
// 	}

// 	// validate user
// 	_, err := repo.GetUserById(r.UserId)
// 	if err != nil {
// 		return fmt.Errorf("No user was found (id: %d)", r.UserId)
// 	}

// 	// validate equip
// 	equip, err := repo.GetEquipmentById(r.EquipId)
// 	if err != nil {
// 		return fmt.Errorf("No equipment was found (id: %d)", r.EquipId)
// 	}

// 	// validate reservation
// 	if r.ReservationId != nil {
// 		_, err = repo.GetEquipmentReservationById(*r.ReservationId)
// 		if err != nil {
// 			return fmt.Errorf("No reservation was found (id: %d)", *r.ReservationId)
// 		}
// 	}

// 	// check equip_sataus and update current history to end
// 	equipUpdate := equipModelToReq(&equip)
// 	if r.EquipStatus == 0 {
// 		equipUpdate.Status = 0
// 		_, err = repo.UpdateEquipment(&equipUpdate)
// 		if err != nil {
// 			return fmt.Errorf("%s", err)
// 		}

// 		currentEquipHistroy, err := repo.GetLatestEquipmentHistoryByEquipId(equip.Id)
// 		if err != nil {
// 			return fmt.Errorf("%s", err)
// 		}
// 		currentEquipHistroyUpdate := equipHistroyModelToReq(&currentEquipHistroy)

// 		_, err = repo.UpdateEquipmentHistory(&currentEquipHistroyUpdate)
// 		if err != nil {
// 			return fmt.Errorf("%s", err)
// 		}
// 	}

// 	// create new history to start
// 	if r.EquipStatus == 1 {
// 		equipUpdate.Status = 1
// 		_, err = repo.UpdateEquipment(&equipUpdate)
// 		if err != nil {
// 			return fmt.Errorf("%s", err)
// 		}
// 		_, err = repo.CreateEquipmentHistory(r)
// 		if err != nil {
// 			return fmt.Errorf("%s", err)
// 		}
// 	}

// 	return nil
// }

// func equipModelToReq(e *menty.Equipment) reqenty.EquipmentUpdateRequest {
// 	return reqenty.EquipmentUpdateRequest{
// 		Id:     e.Id,
// 		Status: *e.Status,
// 		EquipmentRequest: &reqenty.EquipmentRequest{
// 			Name: e.Name,
// 		},
// 	}
// }

// func equipHistroyModelToReq(e *menty.EquipmentHistory) reqenty.EquipmentHistoryUpdateRequest {
// 	return reqenty.EquipmentHistoryUpdateRequest{
// 		Id: e.Id,
// 		EquipmentHistoryRequest: &reqenty.EquipmentHistoryRequest{
// 			EquipId:       e.EquipId,
// 			UserId:        e.UserId,
// 			ReservationId: e.ReservationId,
// 		},
// 	}
// }

// add scheuler for the case user forget to stop using equip
