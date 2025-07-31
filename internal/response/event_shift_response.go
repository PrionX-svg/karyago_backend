package response

import "time"

type EventShiftResponse struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	EventUUID string    `json:"event_uuid"`
	Date      time.Time `json:"date"`
	HourFrom  time.Time `json:"hour_from"`
	HourUntil time.Time `json:"hour_until"`
	Event     struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}
}
