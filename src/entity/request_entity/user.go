package request_entity

type UserRequest struct {
	Authority_id int    `json:"auth_id" binding:"required"`
	Google_id    *int   `json:"google_id default=nil"`
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

type UserUpdateRequest struct {
	Authority_id int    `json:"auth_id"`
	Google_id    *int   `json:"google_id default=nil"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password""`
	Id           int    `json:"id"`
}
