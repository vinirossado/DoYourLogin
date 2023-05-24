package responses

type Response struct {
	StatusCode int `json:"status_code"`
	Data       any `json:"data"`
}
