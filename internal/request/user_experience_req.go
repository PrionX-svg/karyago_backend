package request

import "time"

type UserExperienceReq struct {
	UserUUID  string    `json:"user_uuid" validate:"required"`
	Name      string    `json:"name" validate:"required,max=100"`
	Description  string    `json:"description" validate:"required,max=1000"`
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"omitempty"`
}
