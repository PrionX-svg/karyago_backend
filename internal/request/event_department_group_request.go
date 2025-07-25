package request

type EventDepartmentGroupRequest struct {
	EventUUID       string  `json:"event_uuid" validate:"required,uuid4"`
	Name            string  `json:"name" validate:"required"`
	Description     *string `json:"description,omitempty"`
	ResponsibleUUID *string `json:"responsible_uuid,omitempty"`
}
