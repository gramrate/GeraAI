package responses

type AuthDTO struct {
	Status string `json:"status" example:"success"`
	Token  string `json:"token" example:"your-jwt-token"`
}
