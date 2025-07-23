package request

import "time"

type UserBankReq struct {
	UserUUID  string    `json:"user_uuid" validate:"required"`
	Name      string    `json:"name" validate:"required,max=100"`
	Number    string    `json:"number" validate:"required,max=50"`
	ExpDate   time.Time `json:"exp_date" validate:"required"`
}
