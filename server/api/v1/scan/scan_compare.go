package scan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/scan"
	scanReq "github.com/flipped-aurora/gin-vue-admin/server/model/scan/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ScanApi struct{}

// CreateScan 创建扫码对比
// @Tags Scan
// @Summary 创建扫码对比
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body scan.Scan true "创建扫码对比"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /scan_compare/createScan [post]
func (scan_compareApi *ScanApi) CreateScan(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var scan_compare scan.Scan
	err := c.ShouldBindJSON(&scan_compare)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scan_compareService.CreateScan(ctx, &scan_compare)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteScan 删除扫码对比
// @Tags Scan
// @Summary 删除扫码对比
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body scan.Scan true "删除扫码对比"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /scan_compare/deleteScan [delete]
func (scan_compareApi *ScanApi) DeleteScan(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := scan_compareService.DeleteScan(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteScanByIds 批量删除扫码对比
// @Tags Scan
// @Summary 批量删除扫码对比
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /scan_compare/deleteScanByIds [delete]
func (scan_compareApi *ScanApi) DeleteScanByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := scan_compareService.DeleteScanByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateScan 更新扫码对比
// @Tags Scan
// @Summary 更新扫码对比
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body scan.Scan true "更新扫码对比"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /scan_compare/updateScan [put]
func (scan_compareApi *ScanApi) UpdateScan(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var scan_compare scan.Scan
	err := c.ShouldBindJSON(&scan_compare)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scan_compareService.UpdateScan(ctx, scan_compare)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindScan 用id查询扫码对比
// @Tags Scan
// @Summary 用id查询扫码对比
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询扫码对比"
// @Success 200 {object} response.Response{data=scan.Scan,msg=string} "查询成功"
// @Router /scan_compare/findScan [get]
func (scan_compareApi *ScanApi) FindScan(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rescan_compare, err := scan_compareService.GetScan(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rescan_compare, c)
}

// GetScanList 分页获取扫码对比列表
// @Tags Scan
// @Summary 分页获取扫码对比列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query scanReq.ScanSearch true "分页获取扫码对比列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /scan_compare/getScanList [get]
func (scan_compareApi *ScanApi) GetScanList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo scanReq.ScanSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := scan_compareService.GetScanInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetScanPublic 不需要鉴权的扫码对比接口
// @Tags Scan
// @Summary 不需要鉴权的扫码对比接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /scan_compare/getScanPublic [get]
func (scan_compareApi *ScanApi) GetScanPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	scan_compareService.GetScanPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的扫码对比接口信息",
	}, "获取成功", c)
}

// GetScanInfoPublic 不需要鉴权的扫码对比接口GET
// @Tags Scan
// @Summary 不需要鉴权的扫码对比接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /scan_compare [get]
func (scan_compareApi *ScanApi) GetScanInfoPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	scan_compareService.HandleScanInfoPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的扫码对比接口信息GET",
	}, "获取成功", c)
}

// PostScanInfoPublic 不需要鉴权的扫码对比接口GET
// @Tags Scan
// @Summary 不需要鉴权的扫码对比接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /scan_compare [post]
func (scan_compareApi *ScanApi) PostScanInfoPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	scan_compareService.HandleScanInfoPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的扫码对比接口信息POST",
	}, "获取成功", c)
}
