package service

import (
	"github.com/alice52/archive/bilibili/service/bili"
	"github.com/alice52/archive/bilibili/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	BiliServiceGroup   bili.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
