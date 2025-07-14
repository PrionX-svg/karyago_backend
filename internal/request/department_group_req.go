package request

type DepartmentGroupReq struct {
	CompanyUUID     string `json:"company_uuid" validate:"required"`
	ResponsibleUUID string `json:"responsible_uuid"`
	Name            string `json:"name" validate:"required,max=100"`
	Desc            string `json:"desc" validate:"required,max=255"`
}