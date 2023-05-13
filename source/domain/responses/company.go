package response

type CompanyResponse struct {
	Name string `json:"name" binding:"required" example:"NameOfYourCompany"`
}
