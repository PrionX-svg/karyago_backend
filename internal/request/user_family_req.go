package request

type UserFamilyRequest struct {
	UserUUID string   `json:"user_uuid" validate:"required,uuid4"`
	Status   *bool     `json:"status" validate:"required"`
	Members []struct {
		Name     string `json:"name" validate:"required,max=255"`
		Relation string `json:"relation" validate:"required,max=100"`
		Phone    string `json:"phone" validate:"required,max=20"`
	} `json:"members" validate:"required,dive"`
}

