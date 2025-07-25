package response

type EventDepartmentResponse struct {
	UUID        string  `json:"uuid"`
	GroupUUID   string  `json:"group_uuid"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

