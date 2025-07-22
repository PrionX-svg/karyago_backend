package response

type UserFamilyResponse struct {
	UUID    string   `json:"uuid"`
	Status  bool     `json:"status"`
	Count   uint     `json:"count"`
	Members []string `json:"members"`

	User struct {
		UUID     string `json:"uuid"`
		FullName string `json:"fullname"`
		Email    string `json:"email"`
	} `json:"user"`
}