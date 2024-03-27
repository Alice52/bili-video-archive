package service

type ServiceGroup struct {
	UserUpperTagServiceIn
	UserUpperServiceIn
	UserFavFolderServiceIn
	UserFavServiceIn
	UserLikedServiceIn
	UserCoinedServiceIn
}

var ServiceGroupApp = new(ServiceGroup)

var (
	UserUpperTagService  = ServiceGroupApp.UserUpperTagServiceIn
	UserUpperService     = ServiceGroupApp.UserUpperServiceIn
	UserFavFolderService = ServiceGroupApp.UserFavFolderServiceIn
	UserFavService       = ServiceGroupApp.UserFavServiceIn
	UserLikedService     = ServiceGroupApp.UserLikedServiceIn
	UserCoinedService    = ServiceGroupApp.UserCoinedServiceIn
)
