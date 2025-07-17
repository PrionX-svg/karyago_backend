package response

import "time"

type UserWithEmployeeResponse struct {
	UserUUID     string     `json:"user_uuid"`
	EmployeeUUID string     `json:"employee_uuid"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	FullName     string     `json:"full_name"`
	Email        string     `json:"email"`
	Phone        string     `json:"phone"`
	Gender       *string    `json:"gender"`
	DOB          *time.Time `json:"dob"`
	IsFreelance  bool       `json:"is_freelance"`
	Role         struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	} `json:"role"`
	Branch struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	} `json:"branch"`
	Termination *struct {
		Reason string     `json:"reason"`
		Date   *time.Time `json:"date"`
	} `json:"termination,omitempty"`
	Company *struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	} `json:"company,omitempty"`
}
