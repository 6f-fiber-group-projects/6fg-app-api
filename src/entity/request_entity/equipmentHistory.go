package request_entity

type EquipmentHistoryRequest struct {
	ReservationId *int `json:"reservation_id, default=nil" `
	EquipId       int  `json:"equip_id"`
	EquipStatus   *int `json:"equip_status"` // requiered but 0 cause error https://github.com/gin-gonic/gin/issues/1246
	UserId        int  `json:"user_id" binding:"required"`
}

type EquipmentHistoryUpdateRequest struct {
	*EquipmentHistoryRequest
	Id int
}
