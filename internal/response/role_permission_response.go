package response

type RolePermissionGrouped struct {
	RoleName    string              `json:"role_name"`
	Permissions []PermissionSummary `json:"permissions"`
}

type PermissionSummary struct {
	PermissionName string `json:"permission_name"`
	Label          string `json:"label"`
}
