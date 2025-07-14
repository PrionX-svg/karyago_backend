package response

type DepartmentGroupResponse struct {
	UUID            string `json:"uuid"`
	CompanyUUID     string `json:"company_uuid"`
	ResponsibleUUID string `json:"responsible_uuid,omitempty"`
	Name            string `json:"name"`
	Desc            string `json:"desc"`
}
