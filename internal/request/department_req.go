package request

type DepartmentReq struct {
	DepartmentGroupUUID string `json:"department_group_uuid" validate:"required,uuid4"`
	Name                string `json:"name" validate:"required"`
	Description         string `json:"description"`
	CreatedBy           uint   `json:"created_by"`
	ModifyBy            uint   `json:"modify_by"`
}
