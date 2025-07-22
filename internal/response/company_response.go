package response

type CompanyResponse struct {
	UUID    string `json:"uuid"`
	Logo    string `json:"logo"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	User    struct {
		UUID      string  `json:"uuid"`
		FirstName string  `json:"firstname"`
		LastName  string  `json:"lastname"`
		Role      *string `json:"role,omitempty"`
	} `json:"user"`
}
