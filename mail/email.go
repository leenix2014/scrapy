package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

var from = "liquanlin@liquanlin.tech"
var password = "Leen123"
var smtpServer = "smtp.exmail.qq.com"

func SendPlain(to, subject, body string) error {
	return sendMail(to, subject, body, "")
}

func SendHtml(to, subject, body string) error {
	return sendMail(to, subject, body, "html")
}

func sendMail(to, subject, body, mailType string) error {
	auth := smtp.PlainAuth("", from, password, smtpServer)
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + from + "<" + from + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTos := strings.Split(to, ";")
	err := smtp.SendMail(smtpServer+":25", auth, from, sendTos, msg)
	if err != nil {
		fmt.Println("发送邮件失败!")
		fmt.Println(err)
	} else {
		fmt.Println("发送邮件成功!")
	}
	return err
}
