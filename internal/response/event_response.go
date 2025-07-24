package response

import "time"

type EventResponse struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Company   struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	} `json:"company"`
}
