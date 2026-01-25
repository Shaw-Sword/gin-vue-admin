package service

import (
	"gva/service/example"
	"gva/service/system"
	"gva/service/test_task"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	Test_taskServiceGroup test_task.ServiceGroup
}
