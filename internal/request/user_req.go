package request

type UserRequest struct {
	CompanyID  *uint  `json:"company_id"` 
	FirstName  string `json:"firstname" validate:"required"`
	LastName   string `json:"lastname" validate:"required"`
	Phone      string `json:"phone"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6"`
	Timezone   string `json:"timezone" validate:"required"`
}
