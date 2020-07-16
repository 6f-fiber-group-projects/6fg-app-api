package request_entity

type EquipmentRequest struct {
	Name string `json:"name" binding:"required"`
}

type EquipmentUpdateRequest struct {
	*EquipmentRequest
	Id     int `json:"id" binding:"required"`
	Status int `json:"id`
}
