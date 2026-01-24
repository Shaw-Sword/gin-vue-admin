package test_task

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ BTaskRouter }

var CTaskApi = api.ApiGroupApp.Test_taskApiGroup.BTaskApi
