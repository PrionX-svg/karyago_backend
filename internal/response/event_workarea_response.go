package response

type EventWorkAreaResponse struct {
	UUID      string `json:"uuid"`
	EventUUID string `json:"event_uuid"`
	Name      string `json:"name"`
}
