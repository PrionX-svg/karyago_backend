package response

type EventItemResponse struct {
	UUID      string `json:"uuid"`
	EventUUID string `json:"event_uuid"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}
