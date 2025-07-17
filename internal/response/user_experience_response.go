package response

type UserExperienceResponse struct {
	UUID      string  `json:"uuid"`
	Name      string  `json:"name"`
	Description  string  `json:"description"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date,omitempty"`
	User      struct {
		UUID      string `json:"uuid"`
		FullName string `json:"fullname"`
		Email     string `json:"email"`
	} `json:"user"`
}
