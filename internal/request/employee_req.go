package request

type RehireEmployeeReq struct {
	RoleUUID    string `json:"role_uuid"`
	BranchUUID  string `json:"branch_uuid,omitempty"`
	IsFreelance bool   `json:"is_freelance"`
}
