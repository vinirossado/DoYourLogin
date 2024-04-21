package responses

type AuthorResponse struct {
	ID    uint   `json:"id" example:"1"`
	Name  string `json:"name" binding:"required" example:"NameOfYourCompany"`
	Email string `json:"email" binding:"required" example:"email@email.com"`
}
