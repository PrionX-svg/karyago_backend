package request

import "time"

type UserEmployeeReq struct {
	RoleUUID    string `json:"role_uuid" validate:"required,uuid"`
	CompanyUUID string `json:"company_uuid" validate:"required,uuid"`
	FirstName   string `json:"firstname" validate:"required"`
	LastName    string `json:"lastname" validate:"required"`
	Phone       string `json:"phone"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`

	BranchUUID  string    `json:"branch_uuid" validate:"required,uuid"`
	DOB         time.Time `json:"dob" validate:"required"`
	Gender      string    `json:"gender" validate:"required,oneof=male female"`
	IsFreelance bool      `json:"is_freelance"`
}
