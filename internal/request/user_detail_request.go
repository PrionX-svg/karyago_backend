package request

type UserDetailReq struct {
	UserUUID string `json:"user_uuid" validate:"required,uuid"`
	TaxID    string `json:"tax_id" validate:"required"`    // NPWP
	SocialID string `json:"social_id" validate:"required"` // KTP
}
