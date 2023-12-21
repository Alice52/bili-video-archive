package system

import "github.com/alice52/archive/common/util"

type EmailService struct{}

func (e *EmailService) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = util.EmailTest(subject, body)
	return err
}

func (e *EmailService) SendEmail(to, subject, body string) (err error) {
	err = util.Email(to, subject, body)
	return err
}

func (e *EmailService) SendAttach(to, subject string, att ...string) (err error) {
	err = util.Email(to, subject, "", att...)
	return err
}
