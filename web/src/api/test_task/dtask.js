import service from '@/utils/request'
// @Tags BTask
// @Summary 创建测试任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BTask true "创建测试任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /CTask/createBTask [post]
export const createBTask = (data) => {
  return service({
    url: '/CTask/createBTask',
    method: 'post',
    data
  })
}

// @Tags BTask
// @Summary 删除测试任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BTask true "删除测试任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /CTask/deleteBTask [delete]
export const deleteBTask = (params) => {
  return service({
    url: '/CTask/deleteBTask',
    method: 'delete',
    params
  })
}

// @Tags BTask
// @Summary 批量删除测试任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /CTask/deleteBTask [delete]
export const deleteBTaskByIds = (params) => {
  return service({
    url: '/CTask/deleteBTaskByIds',
    method: 'delete',
    params
  })
}

// @Tags BTask
// @Summary 更新测试任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BTask true "更新测试任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /CTask/updateBTask [put]
export const updateBTask = (data) => {
  return service({
    url: '/CTask/updateBTask',
    method: 'put',
    data
  })
}

// @Tags BTask
// @Summary 用id查询测试任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.BTask true "用id查询测试任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /CTask/findBTask [get]
export const findBTask = (params) => {
  return service({
    url: '/CTask/findBTask',
    method: 'get',
    params
  })
}

// @Tags BTask
// @Summary 分页获取测试任务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /CTask/getBTaskList [get]
export const getBTaskList = (params) => {
  return service({
    url: '/CTask/getBTaskList',
    method: 'get',
    params
  })
}

// @Tags BTask
// @Summary 不需要鉴权的测试任务接口
// @Accept application/json
// @Produce application/json
// @Param data query test_taskReq.BTaskSearch true "分页获取测试任务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /CTask/getBTaskPublic [get]
export const getBTaskPublic = () => {
  return service({
    url: '/CTask/getBTaskPublic',
    method: 'get',
  })
}
// Test_fn_t1 测试方法1
// @Tags BTask
// @Summary 测试方法1
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /CTask/test_fn1 [POST]
export const test_fn1 = () => {
  return service({
    url: '/CTask/test_fn1',
    method: 'POST'
  })
}
