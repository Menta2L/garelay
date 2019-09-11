package main
type Recipient struct {
	Email string `json:"email"`
	Name string `json:"name"`
}
type Message struct {
	Html string  `json:"html"`
	Text string  `json:"text"`
	Subject string  `json:"subject"`
	To []Recipient  `json:"to"`
	FromEmail string  `json:"from_email"`
	FromName string `json:"from_name"`
	Mailclass string `json:"mailclass"`
}
type GAMessage struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Message Message `json:"message"`
}
type GAResponse struct {
	Success int `json:"success"`
	MessageId string `json:"message_id"`
	Error string `json:"error"`
}

type config struct {
	Address      string    `env:"BIND_ADDRESS" envDefault:"0.0.0.0"`
	Port         uint      `env:"BIND_PORT" envDefault:"25"`
	User         string    `env:"GA_USER,required"`
	Password     string    `env:"GA_PASSWORD,required"`
	MailClass    string    `env:"GA_MAIL_CLASS,required"`
	Url          string    `env:"GA_URL,required"`
}