package response

import "time"

type UserWithEmployeeResponse struct {
	UserUUID    string     `json:"user_uuid"`
	FullName    string     `json:"full_name"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	Gender      *string    `json:"gender"`
	DOB         *time.Time `json:"dob"`
	IsFreelance bool       `json:"is_freelance"`
	Role        string     `json:"role"`
	Branch      struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	} `json:"branch"`
	Termination *struct {
		Reason string     `json:"reason"`
		Date   *time.Time `json:"date"`
	} `json:"termination,omitempty"`
}
