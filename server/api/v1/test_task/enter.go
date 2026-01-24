package test_task

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ BTaskApi }

var CTaskService = service.ServiceGroupApp.Test_taskServiceGroup.BTaskService
