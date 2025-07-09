package request

type CompanyReq struct {
	UserUUID string `json:"user_uuid" validate:"required,uuid"`
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone"`
	Logo     string `json:"logo"`
}
