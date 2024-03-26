package service

type ServiceGroup struct {
	UserUpperTagServiceIn
	UserUpperServiceIn
}

var ServiceGroupApp = new(ServiceGroup)

var (
	UserUpperTagService = ServiceGroupApp.UserUpperTagServiceIn
	UserUpperService    = ServiceGroupApp.UserUpperServiceIn
)
