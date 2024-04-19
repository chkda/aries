package gmail

type Config struct {
	SMTPServer string `json:"smtp_server"`
	SMTPPort   string `json:"smtp_port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}
