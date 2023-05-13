package responses

type CompanyResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	ApiToken string `json:"api_token"`
}
