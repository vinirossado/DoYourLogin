package responses

type CompanyResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	APIToken string `json:"api_token"`
}
