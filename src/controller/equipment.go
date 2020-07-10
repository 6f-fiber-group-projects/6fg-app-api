package controller

import (
	bl "6fg-app-api/businessLogic"
	menty "6fg-app-api/entity/model_entity"
	reqenty "6fg-app-api/entity/request_entity"
	resenty "6fg-app-api/entity/response_entity"
	repo "6fg-app-api/repository"
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetEquipments(c *gin.Context) {
	equips, err := repo.GetAllEquipments()
	if err != nil {
		ResponseErrorMessage(c, "#S3AB7J44", "No equipments")
		return
	}

	formatedEquipments := []resenty.EquipmentResponse{}
	for _, equip := range equips {
		formatedEquipments = append(formatedEquipments, formatEquipmentResponse(equip))
	}

	c.JSON(http.StatusOK, gin.H{"message": formatedEquipments})
}

func CreateEquipment(c *gin.Context) {
	equip := reqenty.EquipmentRequest{}
	err := c.ShouldBindJSON(&equip)
	if err != nil {
		ResponseErrorMessage(c, "#NN6O9248", "Bad request")
		return
	}

	_, err = repo.CreateEquipment(&equip)
	if err != nil {
		ResponseErrorMessage(c, "#SINM0QRJ", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

// equipment/:id
func GetEquipmentById(c *gin.Context) {
	equipId, err := strconv.Atoi(c.Param("equipId"))
	if err != nil {
		ResponseErrorMessage(c, "#V60J5UJK", "Equip id shoud be integer")
		return
	}

	equip, err := repo.GetEquipmentById(equipId)
	if err != nil {
		ResponseErrorMessage(c, "#678EZ5VD", "No equipment found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": equip})
}

func UpdateEquipment(c *gin.Context) {
	equip := reqenty.EquipmentUpdateRequest{}
	err := c.ShouldBindJSON(&equip)
	if err != nil {
		ResponseErrorMessage(c, "#W0WAVEGS", "Bad request")
		return
	}

	_, err = repo.UpdateEquipment(&equip)
	if err != nil {
		ResponseErrorMessage(c, "#51IVOGXD", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

func DeleteEquipment(c *gin.Context) {
	equipId, err := strconv.Atoi(c.Param("equipId"))
	if err != nil {
		ResponseErrorMessage(c, "#V60J5UJK", "Equip id shoud be integer")
		return
	}

	_, err = repo.DeleteEquipment(equipId)
	if err != nil {
		ResponseErrorMessage(c, "#AMAO79PX", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

func GetEquipmentQRcode(c *gin.Context) {
	equipId, err := strconv.Atoi(c.Param("equipId"))
	if err != nil {
		ResponseErrorMessage(c, "#68ON4S7N", "Equip id shoud be integer")
		return
	}

	qr, err := bl.GenerateEquipmentQR(equipId)
	if err != nil {
		ResponseErrorMessage(c, "#UB3N0VYD", err.Error())
		return
	}

	c.Data(http.StatusOK, "image/png", qr)
}

func formatEquipmentResponse(e menty.Equipment) resenty.EquipmentResponse {
	return resenty.EquipmentResponse{
		Id:   e.Id,
		Name: e.Name,
	}
}
