package request_entity

type EquipmentRequest struct {
	Name string `json:name`
}

type EquipmentUpdateRequest struct {
	*EquipmentRequest
	Id int `json: id`
}
