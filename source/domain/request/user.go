package request

import "abrigos/source/domain/enumerations"

type UserRequest struct {
	Name                 string             `json:"name" binding:"required" example:"Vinicius Rossado"`
	Password             string             `json:"password" binding:"required" example:"Teste@100"`
	ConfirmationPassword string             `json:"confirmationPassword" binding:"required" example:"Teste@100"`
	Username             string             `json:"username" binding:"required" example:"vinirossado"`
	Role                 enumerations.Roles `json:"role" example:"ADMIN"`
	Email                string             `json:"email" binding:"required" example:"testedeemail@gmail.com"`
	Address              string             `json:"address" binding:"required" example:"Rua dos codigos 522"`
	Phone                string             `json:"phone" binding:"required" example:"+372 0000-0000"`
	About                string             `json:"about"`
	Image                string             `json:"image"`
}
