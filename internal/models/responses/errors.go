package responses

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type ValidationErrorResponse struct {
	Status string            `json:"status"`
	Errors map[string]string `json:"errors"`
}
