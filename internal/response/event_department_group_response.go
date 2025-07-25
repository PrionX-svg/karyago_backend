package response

type EventDepartmentGroupResponse struct {
	UUID        string             `json:"uuid"`
	EventUUID   string             `json:"event_uuid"`
	Name        string             `json:"name"`
	Description *string            `json:"description,omitempty"`
	Responsible *UserBasicResponse `json:"responsible,omitempty"`
}

type UserBasicResponse struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
}
