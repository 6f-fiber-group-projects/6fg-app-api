package businessLogic

import (
	"fmt"
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	repo "github.com/6f-fiber-group-projects/6fg-app-api/repository"
	"github.com/gin-gonic/gin"
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

	if !canReserve(r.StartDate, r.EndDate, r.EquipId) {
		return fmt.Errorf("unable to reserve")
	}

	_, err = repo.CreateEquipmentReservation(r)
	if err != nil {
		return err
	}

	return nil
}

func UpdateEquipmentReservation(c *gin.Context, r *reqenty.EquipmentReservationUpdateRequest) (*menty.EquipmentReservation, error) {
	if !IsAdmin(c) && !IsSameUser(c, r.UserId) {
		return nil, fmt.Errorf("unauthorized")
	}

	rsvn, err := repo.GetEquipmentReservationById(r.Id)
	if err != nil {
		return nil, err
	}

	if !IsAdmin(c) && r.UserId != 0 && r.UserId != rsvn.UserId {
		return nil, fmt.Errorf("unauthorized")
	}

	if !canReserve(r.StartDate, r.EndDate, r.EquipId) {
		return nil, fmt.Errorf("unable to reserve")
	}

	rsvn, err = repo.UpdateEquipmentReservation(r)
	if err != nil {
		return nil, err
	}

	return &rsvn, nil
}

func DeleteEquipmentReservation(c *gin.Context, r *reqenty.EquipmentReservationDeleteRequest) (*menty.EquipmentReservation, error) {
	if !IsAdmin(c) && !IsSameUser(c, r.UserId) {
		return nil, fmt.Errorf("unauthorized")
	}

	rsvn, err := repo.DeleteEquipmentReservation(r.Id)
	if err != nil {
		return nil, err
	}

	return &rsvn, nil
}

func canReserve(st, ed string, equipId int) bool {
	startDate, endDate, _ := repo.ParseTime(st, ed)
	rsvns, _ := repo.GetEquipmentReservationByEquipId(equipId)
	for _, rsvn := range rsvns {
		if startDate.Before(rsvn.StartDate) && endDate.Before(rsvn.StartDate) {
			continue
		}
		if startDate.After(rsvn.EndDate) && endDate.After(rsvn.EndDate) {
			continue
		}
		return false
	}
	return true
}
