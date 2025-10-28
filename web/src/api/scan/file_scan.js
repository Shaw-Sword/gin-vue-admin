import service from '@/utils/request'
// @Tags Scan
// @Summary 创建扫码对比
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Scan true "创建扫码对比"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /scan_compare/createScan [post]
export const createScan = (data) => {
  return service({
    url: '/scan_compare/createScan',
    method: 'post',
    data
  })
}

// @Tags Scan
// @Summary 删除扫码对比
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Scan true "删除扫码对比"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scan_compare/deleteScan [delete]
export const deleteScan = (params) => {
  return service({
    url: '/scan_compare/deleteScan',
    method: 'delete',
    params
  })
}

// @Tags Scan
// @Summary 批量删除扫码对比
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除扫码对比"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scan_compare/deleteScan [delete]
export const deleteScanByIds = (params) => {
  return service({
    url: '/scan_compare/deleteScanByIds',
    method: 'delete',
    params
  })
}

// @Tags Scan
// @Summary 更新扫码对比
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Scan true "更新扫码对比"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /scan_compare/updateScan [put]
export const updateScan = (data) => {
  return service({
    url: '/scan_compare/updateScan',
    method: 'put',
    data
  })
}

// @Tags Scan
// @Summary 用id查询扫码对比
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Scan true "用id查询扫码对比"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /scan_compare/findScan [get]
export const findScan = (params) => {
  return service({
    url: '/scan_compare/findScan',
    method: 'get',
    params
  })
}

// @Tags Scan
// @Summary 分页获取扫码对比列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取扫码对比列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /scan_compare/getScanList [get]
export const getScanList = (params) => {
  return service({
    url: '/scan_compare/getScanList',
    method: 'get',
    params
  })
}

// @Tags Scan
// @Summary 不需要鉴权的扫码对比接口
// @Accept application/json
// @Produce application/json
// @Param data query scanReq.ScanSearch true "分页获取扫码对比列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /scan_compare/getScanPublic [get]
export const getScanPublic = () => {
  return service({
    url: '/scan_compare/getScanPublic',
    method: 'get',
  })
}
