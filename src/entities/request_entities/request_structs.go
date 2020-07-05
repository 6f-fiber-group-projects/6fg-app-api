package request_entities

type UserRequest struct {
	Authority_id int    `json:"auth_id"`
	Google_id    int    `json:"google_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}
