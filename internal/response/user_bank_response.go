package response

type UserBankResponse struct {
	UUID      string  `json:"uuid"`
	Name      string  `json:"name"`
	Number    string  `json:"number"`
	ExpDate   string  `json:"exp_date"`
	User      struct {
		UUID      string `json:"uuid"`
		FullName string `json:"fullname"`
		Email     string `json:"email"`
	} `json:"user"`
}
