package response_entities

type UserResponse struct {
	Id           int    `json:"id"`
	Authority_id int    `json:"auth_id"`
	Google_id    int    `json:"google_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
}
