package request

type EventShiftReq struct {
	EventUUID string `json:"event_uuid" validate:"required,uuid"`
	Name 	string `json:"name" validate:"required"`
	Date      string `json:"date" validate:"required"`
	HourFrom  string `json:"hour_from" validate:"required"`
	HourUntil string `json:"hour_until" validate:"required"`
}
