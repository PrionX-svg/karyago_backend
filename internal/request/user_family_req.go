package request

type UserFamilyRequest struct {
	UserUUID string   `json:"user_uuid" validate:"required,uuid4"`
	Status   bool     `json:"status" validate:"required"`
	Count    uint     `json:"count" validate:"required"`
	Member   []string `json:"member" validate:"required,dive,required"`
}
