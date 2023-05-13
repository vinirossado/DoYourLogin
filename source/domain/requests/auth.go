package request

type Auth struct {
	Email    string `json:"email" binding:"required" example:"teste@gmail.com"`
	Password string `json:"password" binding:"required" example:"teste@100"`
}
