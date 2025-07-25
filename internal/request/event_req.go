package request

import "time"

type EventReq struct {
	CompanyUUID string      `json:"company_uuid" validate:"required"`
	Name        string    `json:"name" validate:"required,max=100"`
	Photo       *string   `json:"photo" validate:"omitempty,max=255"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"omitempty"`
}
