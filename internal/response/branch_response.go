package response

type BranchResponse struct {
	UUID    string  `json:"uuid"`
	Image   *string `json:"image"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Company struct {
		UUID    string `json:"uuid"`
		Logo    string `json:"logo"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
	} `json:"company"`
}
