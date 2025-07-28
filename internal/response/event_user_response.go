package response

type EventUser struct {
	UUID      string  `json:"uuid"`
	User      struct {
		UUID      string `json:"uuid"`
		FullName string `json:"fullname"`
		Email     string `json:"email"`
	} `json:"user"`
	Event struct {
		UUID        string `json:"uuid"`
		Name        string `json:"name"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
	}
}
