package response

type EventUserResponse struct {
	UUID  string `json:"uuid"`
	User  UserSummary   `json:"user"`
	Event EventSummary  `json:"event"`
}

type UserSummary struct {
	UUID     string `json:"uuid"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type EventSummary struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type EventUserGroupedByUser struct {
	User   UserSummary     `json:"user"`
	Events []EventWithUUID `json:"events"`
}

type EventWithUUID struct {
	EventUserUUID string `json:"uuid"`
	UUID          string `json:"event_uuid"`
	Name          string `json:"name"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
}

type EventUserGroupedByEvent struct {
	Event EventSummary     `json:"event"`
	Users []UserWithUUID   `json:"users"`
}

type UserWithUUID struct {
	EventUserUUID string `json:"uuid"`
	UUID          string `json:"user_uuid"`
	FullName      string `json:"fullname"`
	Email         string `json:"email"`
}