package controller

import (
	bl "github.com/6f-fiber-group-projects/6fg-app-api/businessLogic"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// equipment
func GetEquipments(c *gin.Context) {
	equips, err := bl.GetEquipments(c)
	if err != nil {
		ResponseServerErrorMessage(c, "#S3AB7J44", "No equipments")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": equips})
}

func CreateEquipment(c *gin.Context) {
	equip := reqenty.EquipmentRequest{}
	err := c.ShouldBindJSON(&equip)
	if err != nil {
		ResponseServerErrorMessage(c, "#NN6O9248", "Bad request")
		return
	}

	_, err = bl.CreateEquipment(c, &equip)
	if err != nil {
		ResponseServerErrorMessage(c, "#SINM0QRJ", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

// equipment/:id
func GetEquipmentById(c *gin.Context) {
	equipId, err := strconv.Atoi(c.Param("equipId"))
	if err != nil {
		ResponseServerErrorMessage(c, "#V60J5UJK", "Equip id shoud be integer")
		return
	}

	equip, err := bl.GetEquipmentById(c, equipId)
	if err != nil {
		ResponseServerErrorMessage(c, "#678EZ5VD", "No equipment found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": equip})
}

func UpdateEquipment(c *gin.Context) {
	equipId, err := strconv.Atoi(c.Param("equipId"))
	if err != nil {
		ResponseServerErrorMessage(c, "#", "Equip id shoud be integer")
		return
	}

	equip := reqenty.EquipmentUpdateRequest{}
	err = c.ShouldBindJSON(&equip)
	if err != nil {
		ResponseServerErrorMessage(c, "#W0WAVEGS", "Bad request")
		return
	}
	equip.Id = equipId

	_, err = bl.UpdateEquipment(c, &equip)
	if err != nil {
		ResponseServerErrorMessage(c, "#51IVOGXD", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

func DeleteEquipment(c *gin.Context) {
	equipId, err := strconv.Atoi(c.Param("equipId"))
	if err != nil {
		ResponseServerErrorMessage(c, "#V60J5UJK", "Equip id shoud be integer")
		return
	}

	_, err = bl.DeleteEquipment(c, equipId)
	if err != nil {
		ResponseServerErrorMessage(c, "#AMAO79PX", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

// equipment/:equipId/status

func UpdateEquipmentStatus(c *gin.Context) {
	equipId, err := strconv.Atoi(c.Param("equipId"))
	if err != nil {
		ResponseServerErrorMessage(c, "#", "Equip id shoud be integer")
		return
	}

	equipHistory := reqenty.EquipmentHistoryRequest{}
	err = c.ShouldBindJSON(&equipHistory)
	if err != nil {
		ResponseServerErrorMessage(c, "#", "Bad request")
		return
	}
	equipHistory.EquipId = equipId

	err = bl.UpdateEquipmentStatus(c, &equipHistory)
	if err != nil {
		ResponseServerErrorMessage(c, "#", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}
