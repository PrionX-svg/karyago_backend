package request

import "time"

type ShiftRequest struct {
	CompanyUUID string    `json:"company_uuid" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	HourFrom    time.Time `json:"hour_from" validate:"required"`
	HourUntil   time.Time `json:"hour_until" validate:"required"`
	CreatedBy   string    `json:"-"`
	ModifiedBy  string    `json:"-"`
}
