package request

type EventUserReq struct {
	UserUUID string `json:"user_uuid" validate:"required,uuid4"`
	EventUUID string `json:"event_uuid" validate:"required,uuid4"`
}
