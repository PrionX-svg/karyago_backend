package request

type CompanyReq struct {
	UserId  uint   `json:"user_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Phone   string `json:"phone"`
	Logo    string `json:"logo"`
}
