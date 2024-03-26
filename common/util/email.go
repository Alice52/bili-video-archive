package util

import (
	"crypto/tls"
	"fmt"
	"github.com/alice52/archive/common/global"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

func Email(to, subject string, body string, att ...string) error {
	return send(strings.Split(to, ","), subject, body, att...)
}

func ErrorToEmail(subject string, body string) error {
	to := strings.Split(global.CONFIG.Email.To, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

func EmailTest(subject string, body string) error {
	to := []string{global.CONFIG.Email.To}
	return send(to, subject, body)
}

func send(to []string, subject string, body string, attFileName ...string) error {
	es := global.CONFIG.Email
	from := es.From
	nickname := es.Nickname
	secret := es.Secret
	host := es.Host
	port := es.Port
	isSSL := es.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)

	for _, ap := range attFileName {
		if _, err := e.AttachFile(ap); err != nil {
			fmt.Println("Error attaching file:", err)
			return err
		}
	}

	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host}) //nolint:gosec
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}

func SendEmail(to, subject, body string) (err error) {
	return Email(to, subject, body)
}

func SendAttach(to, subject string, att ...string) (err error) {
	return Email(to, subject, "", att...)
}
