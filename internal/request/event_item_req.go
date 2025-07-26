package request

type EventItemRequest struct {
	UserUUID  string `json:"user_uuid" validate:"required,uuid4"`
	EventUUID string `json:"event_uuid" validate:"required,uuid4"`
	Name      string `json:"name" validate:"required"`
	Type      string `json:"type" validate:"required,oneof=event date shift"`
}
