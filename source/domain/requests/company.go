package requests

type CompanyRequest struct {
	Name     string `json:"name" binding:"required" example:"NameOfYourCompany"`
	Username string `json:"username" binding:"required" example:"yourUsername"`
	Email    string `json:"email" binding:"required" example:"email@email.com"`
	Password string `json:"password" binding:"required" example:"abc*****"`
}
