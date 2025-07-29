package request

type BranchReq struct {
	CompanyUUID string  `json:"company_uuid" validate:"required,uuid"`
	Name        string  `json:"name" validate:"required"`
	Address     string  `json:"address" validate:"required"`
	Email       string  `json:"email" validate:"required,email"`
	Phone       string  `json:"phone"`
	Image       *string `json:"image"`
}
