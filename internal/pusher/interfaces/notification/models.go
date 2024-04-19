package notification

type Request struct {
	NotificationId      string `json:"notification_id"`
	NotificationSubject string `json:"notification_subject"`
	NotificationBody    string `json:"notification_body"`
	UserId              string `json:"user_id"`
	UserDevice          string `json:"user_device"`
	NotificationMethod  string `json:"notification_method"`
	Group               string `json:"group"`
}

type Response struct {
	Success bool    `json:"success"`
	Message *string `json:"message"`
}
