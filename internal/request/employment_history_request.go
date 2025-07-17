package request

import "time"

type EmploymentHistoryRequest struct {
	EmployeeUUID string     `json:"employee_uuid" validate:"required,uuid4"`
	CompanyUUID  string     `json:"company_uuid,omitempty"` // optional
	BranchUUID   string     `json:"branch_uuid,omitempty"`  // optional
	RoleUUID     *string    `json:"role_uuid,omitempty" validate:"uuid4"`
	Position     string     `json:"position" validate:"required"`
	IsPresent    bool       `json:"is_present"`
	StartDate    time.Time  `json:"start_date" validate:"required"`
	EndDate      *time.Time `json:"end_date,omitempty"` // optional if IsPresent = true
	Notes        *string    `json:"notes,omitempty"`
}
