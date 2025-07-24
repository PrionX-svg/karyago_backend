package response

type UserMiniResponse struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DepartmentGroupMiniResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type DepartmentResponse struct {
	UUID        string                      `json:"uuid"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Group       DepartmentGroupMiniResponse `json:"department_group"`
	Employees   []UserMiniResponse          `json:"employees"`
}
