package test_task

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/test_task"
    test_taskReq "github.com/flipped-aurora/gin-vue-admin/server/model/test_task/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type BTaskApi struct {}



// CreateBTask 创建测试任务
// @Tags BTask
// @Summary 创建测试任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body test_task.BTask true "创建测试任务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /CTask/createBTask [post]
func (CTaskApi *BTaskApi) CreateBTask(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var CTask test_task.BTask
	err := c.ShouldBindJSON(&CTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = CTaskService.CreateBTask(ctx,&CTask)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteBTask 删除测试任务
// @Tags BTask
// @Summary 删除测试任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body test_task.BTask true "删除测试任务"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /CTask/deleteBTask [delete]
func (CTaskApi *BTaskApi) DeleteBTask(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := CTaskService.DeleteBTask(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBTaskByIds 批量删除测试任务
// @Tags BTask
// @Summary 批量删除测试任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /CTask/deleteBTaskByIds [delete]
func (CTaskApi *BTaskApi) DeleteBTaskByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := CTaskService.DeleteBTaskByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBTask 更新测试任务
// @Tags BTask
// @Summary 更新测试任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body test_task.BTask true "更新测试任务"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /CTask/updateBTask [put]
func (CTaskApi *BTaskApi) UpdateBTask(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var CTask test_task.BTask
	err := c.ShouldBindJSON(&CTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = CTaskService.UpdateBTask(ctx,CTask)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBTask 用id查询测试任务
// @Tags BTask
// @Summary 用id查询测试任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询测试任务"
// @Success 200 {object} response.Response{data=test_task.BTask,msg=string} "查询成功"
// @Router /CTask/findBTask [get]
func (CTaskApi *BTaskApi) FindBTask(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reCTask, err := CTaskService.GetBTask(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reCTask, c)
}
// GetBTaskList 分页获取测试任务列表
// @Tags BTask
// @Summary 分页获取测试任务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query test_taskReq.BTaskSearch true "分页获取测试任务列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /CTask/getBTaskList [get]
func (CTaskApi *BTaskApi) GetBTaskList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo test_taskReq.BTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := CTaskService.GetBTaskInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetBTaskPublic 不需要鉴权的测试任务接口
// @Tags BTask
// @Summary 不需要鉴权的测试任务接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /CTask/getBTaskPublic [get]
func (CTaskApi *BTaskApi) GetBTaskPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    CTaskService.GetBTaskPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的测试任务接口信息",
    }, "获取成功", c)
}
// Test_fn_t1 测试方法1
// @Tags BTask
// @Summary 测试方法1
// @Accept application/json
// @Produce application/json
// @Param data query test_taskReq.BTaskSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /CTask/test_fn1 [POST]
func (CTaskApi *BTaskApi)Test_fn_t1(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()
    // 请添加自己的业务逻辑
    err := CTaskService.Test_fn_t1(ctx)
    if err != nil {
        global.GVA_LOG.Error("失败!", zap.Error(err))
   		response.FailWithMessage("失败", c)
   		return
   	}
   	response.OkWithData("返回数据",c)
}


