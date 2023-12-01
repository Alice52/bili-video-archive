package util

import (
	"crypto/tls"
	"fmt"
	"github.com/alice52/archive/common/global"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

func Email(To, subject string, body string) error {
	to := strings.Split(To, ",")
	return send(to, subject, body)
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

func send(to []string, subject string, body string) error {
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
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host}) //nolint:gosec
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
