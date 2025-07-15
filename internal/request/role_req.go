package request

type CreateRoleRequest struct {
	Name        string `json:"name"`
	CompanyUUID string `json:"company_uuid"`
}
