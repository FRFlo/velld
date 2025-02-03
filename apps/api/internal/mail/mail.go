package mail

import (
    "fmt"
    "net/smtp"
)

type SMTPConfig struct {
    Host     string
    Port     int
    Username string
    Password string
}

type Message struct {
    From    string
    To      string
    Subject string
    Body    string
}

func SendEmail(config *SMTPConfig, msg *Message) error {
    auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
    
    emailMsg := fmt.Sprintf("From: %s\r\n"+
        "To: %s\r\n"+
        "Subject: %s\r\n"+
        "\r\n"+
        "%s\r\n", msg.From, msg.To, msg.Subject, msg.Body)

    addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
    
    return smtp.SendMail(
        addr,
        auth,
        msg.From,
        []string{msg.To},
        []byte(emailMsg),
    )
}
