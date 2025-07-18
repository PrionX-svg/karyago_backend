package response

import "time"

type UserWithEmployeeAndHistoryResponse struct {
	UserUUID     string                      `json:"user_uuid"`
	EmployeeUUID string                      `json:"employee_uuid"`
	FirstName    string                      `json:"first_name"`
	LastName     string                      `json:"last_name"`
	FullName     string                      `json:"full_name"`
	Email        string                      `json:"email"`
	Phone        string                      `json:"phone"`
	Gender       *string                     `json:"gender,omitempty"`
	DOB          *time.Time                  `json:"dob,omitempty"`
	IsFreelance  bool                        `json:"is_freelance"`
	Role         *RoleSimpleResponse         `json:"role,omitempty"`
	Branch       *BranchSimpleResponse       `json:"branch,omitempty"`
	Company      *CompanySimpleResponse      `json:"company,omitempty"`
	Histories    []EmploymentHistoryResponse `json:"employment_histories,omitempty"`
}

type RoleSimpleResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type BranchSimpleResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type CompanySimpleResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
