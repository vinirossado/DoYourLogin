package request

type CompanyRequest struct {
	Name string `json:"name" binding:"required" example:"NameOfYourCompany"`
}
