package request

import "time"

type EventShiftReq struct {
	EventUUID string `json:"event_uuid" validate:"required,uuid"`
	Name 	string `json:"name" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	HourFrom  time.Time `json:"hour_from" validate:"required"`
	HourUntil time.Time `json:"hour_until" validate:"required"`
}
