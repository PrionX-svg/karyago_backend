package request

type EventDepartmentRequest struct {
	UserUUID    string  `json:"user_uuid" validate:"required,uuid4"`
	GroupUUID   string  `json:"group_uuid" validate:"required,uuid4"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
}

