package service

type ServiceGroup struct {
	EmailServiceIn
}

var ServiceGroupApp = new(ServiceGroup)

var (
	EmailService = ServiceGroupApp.EmailServiceIn
)
