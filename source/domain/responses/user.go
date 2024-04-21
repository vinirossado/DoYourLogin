package responses

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address,omitempty"`
	Phone    string `json:"phone,omitempty"`
	About    string `json:"about,omitempty"`
	Image    string `json:"image,omitempty"`
}
