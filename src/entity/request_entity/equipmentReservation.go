package request_entity

type EquipmentReservationRequest struct {
	EquipId   int    `json:"equipId" binding:"required"`
	UserId    int    `json:"userId" binding:"required"`
	StartDate string `json:"startDate" binding:"required"`
	EndDate   string `json:"endDate" binding:"required"`
}

type EquipmentReservationUpdateRequest struct {
	*EquipmentReservationRequest
	Id int `json:"id"`
}

type EquipmentReservationDeleteRequest struct {
	Id int `json:"id"`
}
