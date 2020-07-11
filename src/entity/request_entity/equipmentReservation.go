package request_entity

type EquipmentReservationRequest struct {
	EquipId   int    `json:"equip_id"`
	UserId    int    `json:"user_id" binding:"required"`
	StartDate string `json:"start_date"  binding:"required"`
	EndDate   string `json:"end_date"  binding:"required"`
}

type EquipmentReservationUpdateRequest struct {
	*EquipmentReservationRequest
	Id int `json:"id" binding:"required"`
}

type EquipmentReservationDeleteRequest struct {
	Id int `json:"id" binding:"required"`
}
