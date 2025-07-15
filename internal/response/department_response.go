package response

type DepartmentResponse struct {
	UUID                string `json:"uuid"`
	DepartmentGroupUUID string `json:"department_group_uuid"`
	Name                string `json:"name"`
	Description         string `json:"description"`
}
