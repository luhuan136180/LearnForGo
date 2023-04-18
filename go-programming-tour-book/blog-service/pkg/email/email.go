package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

func (e *Email) SendMail(to []string, subject, body string) error {
	//NewMessage创建一个新消息。默认情况下，它使用UTF-8和引用可打印编码。
	m := gomail.NewMessage()
	//SetHeader为给定的报头字段设置值。
	//m.Header 是一个哈希map
	m.SetHeader("Form", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	//NewDialer返回一个新的SMTP Dialer。指定的参数用于连接SMTP服务器。
	//这段代码是用来创建一个Dialer对象，这个对象用来连接SMTP服务器并发送邮件。
	//其中，e.Host和e.Port分别指定SMTP服务器的地址和端口号，e.UserName和e.Password是用于SMTP身份验证的用户名和密码。
	//使用这个Dialer对象可以通过调用其Dial方法来建立到SMTP服务器的连接，然后可以发送邮件。
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}
