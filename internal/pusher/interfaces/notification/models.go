package notification

type Request struct {
	EventId             string `json:"event_id"`
	NotificationId      string `json:"notification_id"`
	NotificationSubject string `json:"notification_subject"`
	NotificationBody    string `json:"notification_body"`
	UserId              string `json:"user_id"`
	UserDevice          string `json:"user_device"`
	NotificationType    string `json:"notification_type"`
	BU                  string `json:"bu"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
