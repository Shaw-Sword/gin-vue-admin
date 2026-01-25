package router

import (
	"gva/router/example"
	"gva/router/system"
	"gva/router/test_task"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Test_task test_task.RouterGroup
}
