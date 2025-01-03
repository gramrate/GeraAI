package responses

type GeneratedTaskResponse struct {
	Status        string `json:"status"`
	GeneratedText string `json:"generated_text"`
}

type GeneratedAnswerResponse struct {
	Status        string `json:"status"`
	GeneratedText string `json:"generated_text"`
}
