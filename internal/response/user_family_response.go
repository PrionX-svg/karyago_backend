package response

type UserFamilyResponse struct {
	Status bool `json:"status"`
	Count  uint `json:"count"`

	Members []struct {
		UUID     string `json:"uuid"`
		Name     string `json:"name"`
		Relation string `json:"relation"`
		Phone    string `json:"phone"`
	} `json:"members"`

	User struct {
		UUID     string `json:"uuid"`
		FullName string `json:"fullname"`
		Email    string `json:"email"`
	} `json:"user"`
}
