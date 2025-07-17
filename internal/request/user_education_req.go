package request

import "time"

type UserEducationReq struct {
	UserUUID  string    `json:"user_uuid" validate:"required"`
	Name      string    `json:"name" validate:"required,max=100"`
	Location  string    `json:"location" validate:"required,max=100"`
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"omitempty"`
	Grade     float64   `json:"grade" validate:"required,gte=0,lte=100"`
}
