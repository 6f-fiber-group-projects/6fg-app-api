package request_entity

type EquipmentHistoryRequest struct {
	ReservationId *int `json:"rsvnId, default=nil" `
	EquipId       int  `json:"equipId"`
	EquipStatus   *int `json:"equipStatus"` // requiered but 0 cause error https://github.com/gin-gonic/gin/issues/1246
	UserId        int  `json:"userId" binding:"required"`
}

type EquipmentHistoryUpdateRequest struct {
	*EquipmentHistoryRequest
	Id int
}
