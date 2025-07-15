package request

type UserRequest struct {
	FirstName   string `json:"firstname" validate:"required"`
	LastName    string `json:"lastname" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	Phone       string `json:"phone"`
	Timezone    string `json:"timezone"`
	IsFreelance bool   `json:"is_freelance"` // default: false
}
