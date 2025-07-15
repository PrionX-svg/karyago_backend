package request

import "time"

type UserEmployeeReq struct {
	FirstName   string    `json:"firstname" validate:"required"`
	LastName    string    `json:"lastname" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	Password    string    `json:"password" validate:"required"`
	Phone       string    `json:"phone"`
	DOB         time.Time `json:"dob"`
	Gender      string    `json:"gender"` // "male"/"female"
	IsFreelance bool      `json:"is_freelance"`
	RoleUUID    string    `json:"role_uuid" validate:"required"`
	CompanyUUID string    `json:"company_uuid" validate:"required"` // ✅ wajib
	BranchUUID  string    `json:"branch_uuid"`                      // optional
}
