package request

type EventWorkAreaRequest struct {
	UserUUID  string `json:"user_uuid"`
	EventUUID string `json:"event_uuid"`
	Name      string `json:"name"`
}
