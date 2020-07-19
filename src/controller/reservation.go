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
	rsvnId, err := strconv.Atoi(c.Param("rsvnId"))
	if err != nil {
		ResponseErrorMessage(c, "#", "Reservation id shoud be integer")
		return
	}

	rsvn := reqenty.EquipmentReservationUpdateRequest{}
	err = c.ShouldBindJSON(&rsvn)
	if err != nil {
		ResponseErrorMessage(c, "#Y3VE6O9R", "Bad request")
		return
	}
	rsvn.Id = rsvnId

	_, err = repo.UpdateEquipmentReservation(&rsvn)
	if err != nil {
		ResponseErrorMessage(c, "#WVNAN2JV", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

func DeleteEquipmentReservation(c *gin.Context) {
	rsvnId, err := strconv.Atoi(c.Param("rsvnId"))
	if err != nil {
		ResponseErrorMessage(c, "#I7NB8UY2", "Reservation id shoud be integer")
		return
	}

	_, err = repo.DeleteEquipmentReservation(rsvnId)
	if err != nil {
		ResponseErrorMessage(c, "#37MWJFMZ", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

func GetEquipmentReservation(c *gin.Context) {
	rsvnId, err := strconv.Atoi(c.Param("rsvnId"))
	if err != nil {
		ResponseErrorMessage(c, "#OHGXQ7XW", "Reservation id shoud be integer")
		return
	}

	rsvn, err := repo.GetEquipmentReservationById(rsvnId)
	if err != nil {
		ResponseErrorMessage(c, "#678EZ5VD", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": rsvn})
}

func CreateEquipmentReservation(c *gin.Context) {
	rsvn := reqenty.EquipmentReservationRequest{}
	err := c.ShouldBindJSON(&rsvn)
	if err != nil {
		ResponseErrorMessage(c, "#UB3N0VYD", err.Error())
		return
	}

	err = bl.ReserveEquip(&rsvn)
	if err != nil {
		ResponseErrorMessage(c, "#0R8INES3", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}
