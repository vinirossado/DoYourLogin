package responses

import (
	"abrigos/source/domain/enumerations"
)

type UserResponse struct {
	ID       uint               `json:"id"`
	Name     string             `json:"name"`
	Username string             `json:"username"`
	Role     enumerations.Roles `json:"role"`
	Email    string             `json:"email"`
	Address  string             `json:"address"`
	Phone    string             `json:"phone"`
	About    string             `json:"about"`
	Image    string             `json:"image"`
}
