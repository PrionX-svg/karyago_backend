package response

type UserEducationResponse struct {
	UUID      string  `json:"uuid"`
	Name      string  `json:"name"`
	Location  string  `json:"location"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date,omitempty"`
	Grade     float64 `json:"grade"`
	User      struct {
		UUID      string `json:"uuid"`
		FullName string `json:"fullname"`
		Email     string `json:"email"`
	} `json:"user"`
}
