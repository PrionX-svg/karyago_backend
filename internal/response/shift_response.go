package response

import "time"

type ShiftResponse struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	HourFrom  time.Time `json:"hour_from"`
	HourUntil time.Time `json:"hour_until"`
}
