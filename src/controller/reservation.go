package controller

import (
	bl "6fg-app-api/businessLogic"
	reqenty "6fg-app-api/entity/request_entity"
	repo "6fg-app-api/repository"
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateEquipmentReservation(c *gin.Context) {
	rsvn := reqenty.EquipmentReservationUpdateRequest{}
	err := c.ShouldBindJSON(&rsvn)
	if err != nil {
		ResponseErrorMessage(c, "#Y3VE6O9R", "Bad request")
		return
	}

	_, err = repo.UpdateEquipmentReservation(&rsvn)
	if err != nil {
		ResponseErrorMessage(c, "#WVNAN2JV", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

func DeleteEquipmentReservation(c *gin.Context) {
	rsvn := reqenty.EquipmentReservationDeleteRequest{}
	err := c.ShouldBindJSON(&rsvn)
	if err != nil {
		ResponseErrorMessage(c, "#I7NB8UY2", "Bad request")
		return
	}

	_, err = repo.DeleteEquipmentReservation(&rsvn)
	if err != nil {
		ResponseErrorMessage(c, "#37MWJFMZ", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

func GetEquipmentReservation(c *gin.Context) {
	equipId, err := strconv.Atoi(c.Param("equipId"))
	if err != nil {
		ResponseErrorMessage(c, "#OHGXQ7XW", "Equip id shoud be integer")
		return
	}

	rsvns, err := repo.GetEquipmentReservationByEquipId(equipId)
	if err != nil {
		ResponseErrorMessage(c, "#678EZ5VD", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": rsvns})
}

func CreateEquipmentReservation(c *gin.Context) {
	equipId, err := strconv.Atoi(c.Param("equipId"))
	if err != nil {
		ResponseErrorMessage(c, "#XOQ0B093", "Equip id shoud be integer")
		return
	}

	rsvn := reqenty.EquipmentReservationRequest{}
	err = c.ShouldBindJSON(&rsvn)
	if err != nil {
		ResponseErrorMessage(c, "#UB3N0VYD", err.Error())
		return
	}
	rsvn.EquipId = equipId

	err = bl.ReserveEquip(&rsvn)
	if err != nil {
		ResponseErrorMessage(c, "#0R8INES3", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}
