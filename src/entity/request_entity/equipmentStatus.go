package request_entity

type EquipmentStatusUpdateRequest struct {
	EquipId int
	UserId  *int
	Status  *int
}
