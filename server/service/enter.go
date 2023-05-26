package service

import (
	"github.com/VINDA-98/Tenacity-Solutions/server/service/example"
	"github.com/VINDA-98/Tenacity-Solutions/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
