package request

type ChatbotRequest struct {
	Model   string `json:"model"`
	Message string `json:"message" validate:"required"`
}
