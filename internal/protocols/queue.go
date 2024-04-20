package protocols

type Message struct {
	EventId             string `json:"event_id"`
	NotificationId      string `json:"notification_id"`
	NotificationSubject string `json:"notification_subject"`
	NotificationBody    string `json:"notification_body"`
	UserId              string `json:"user_id"`
	UserDevice          string `json:"user_device"`
	NotificationType    string `json:"notification_type"`
	BU                  string `json:"bu"`
}

const NOTFICATION_EVENTS_QUEUE = "notification_events_v1"
