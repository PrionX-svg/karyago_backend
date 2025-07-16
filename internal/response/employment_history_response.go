package response

import "time"

type EmploymentHistoryResponse struct {
	UUID      string                 `json:"uuid"`
	Employee  SimpleEmployeeResponse `json:"employee"`
	Company   *SimpleCompanyResponse `json:"company,omitempty"`
	Branch    *SimpleBranchResponse  `json:"branch,omitempty"`
	Role      SimpleRoleResponse     `json:"role"`
	Position  string                 `json:"position"`
	IsPresent bool                   `json:"is_present"`
	StartDate time.Time              `json:"start_date"`
	EndDate   *time.Time             `json:"end_date,omitempty"`
	Notes     *string                `json:"notes,omitempty"`
}

type SimpleEmployeeResponse struct {
	UUID     string `json:"uuid"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type SimpleCompanyResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type SimpleBranchResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type SimpleRoleResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
