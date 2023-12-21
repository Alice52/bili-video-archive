package system

type ServiceGroup struct {
	EmailService
}

var SystemService = new(ServiceGroup)
