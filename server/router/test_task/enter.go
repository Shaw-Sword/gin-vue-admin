package test_task

import api "gva/api/v1"

type RouterGroup struct{ BTaskRouter }

var CTaskApi = api.ApiGroupApp.Test_taskApiGroup.BTaskApi
