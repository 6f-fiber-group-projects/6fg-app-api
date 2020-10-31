package businessLogic

import (
	"fmt"
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	resenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/response_entity"
	repo "github.com/6f-fiber-group-projects/6fg-app-api/repository"
	"github.com/gin-gonic/gin"
)

func GetEquipments(c *gin.Context) (*[]resenty.EquipmentResponse, error) {
	equips, err := repo.GetAllEquipments()
	if err != nil {
		return nil, err
	}

	equipRes := []resenty.EquipmentResponse{}
	for _, equip := range equips {
		equipStatus, err := repo.GetEquipmentStatusByEquipId(equip.Id)
		if err != nil {
			continue
		}
		equipRes = append(equipRes, formatEquipmentResponse(&equip, &equipStatus))
	}

	return &equipRes, nil
}

func GetEquipmentById(c *gin.Context, id int) (*resenty.EquipmentResponse, error) {
	equip, err := repo.GetEquipmentById(id)
	if err != nil {
		return nil, err
	}

	equipStatus, _ := repo.GetEquipmentStatusByEquipId(equip.Id)
	equipRes := formatEquipmentResponse(&equip, &equipStatus)

	return &equipRes, nil
}

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

func UpdateEquipmentStatus(c *gin.Context, r *reqenty.EquipmentHistoryRequest) error {
	if !IsAdmin(c) && !IsSameUser(c, r.UserId) {
		return fmt.Errorf("unauthorized")
	}

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
	if *r.EquipStatus == 0 {
		equipStausUpdate := equipStatusUpdateReq(r)
		_, err = repo.UpdateEquipmentStatus(&equipStausUpdate)
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
	if *r.EquipStatus == 1 {
		equipStausUpdate := equipStatusUpdateReq(r)
		_, err = repo.UpdateEquipmentStatus(&equipStausUpdate)
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

func equipStatusUpdateReq(r *reqenty.EquipmentHistoryRequest) reqenty.EquipmentStatusUpdateRequest {
	return reqenty.EquipmentStatusUpdateRequest{
		EquipId: r.EquipId,
		UserId:  &r.UserId,
		Status:  r.EquipStatus,
	}
}

func formatEquipmentResponse(e *menty.Equipment, es *menty.EquipmentStatus) resenty.EquipmentResponse {
	return resenty.EquipmentResponse{
		Id:     e.Id,
		Name:   e.Name,
		Status: *es.Status,
		UserId: *es.UserId,
	}
}

// add scheuler for the case user forget to stop using equip
