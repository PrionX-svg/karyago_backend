package response

type DepartmentGroupResponse struct {
	UUID        string               `json:"uuid"`
	CompanyUUID string               `json:"company_uuid"`
	Responsible *ResponsibleResponse `json:"responsible,omitempty"`
	Name        string               `json:"name"`
	Desc        string               `json:"desc"`
}

type ResponsibleResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
