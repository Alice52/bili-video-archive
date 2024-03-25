package service

import "github.com/alice52/archive/common/util"

type EmailServiceIn struct{}

func (e *EmailServiceIn) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = util.EmailTest(subject, body)
	return err
}

func (e *EmailServiceIn) SendEmail(to, subject, body string) (err error) {
	return util.Email(to, subject, body)
}

func (e *EmailServiceIn) SendAttach(to, subject string, att ...string) (err error) {
	return util.Email(to, subject, "", att...)
}
