package requests

type AuthorRequest struct {
	Name  string `json:"name" binding:"required" example:"NameOfYourAuthor"`
	Email string `json:"email" binding:"required" example:"email@email.com"`
}
