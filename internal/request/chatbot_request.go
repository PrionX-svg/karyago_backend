package request

type ChatbotRequest struct {
	Model   string `json:"model" validate:"required"`
	Message string `json:"message" validate:"required"`
}
