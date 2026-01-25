package v1

import (
	"gva/api/v1/example"
	"gva/api/v1/system"
	"gva/api/v1/test_task"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	Test_taskApiGroup test_task.ApiGroup
}
