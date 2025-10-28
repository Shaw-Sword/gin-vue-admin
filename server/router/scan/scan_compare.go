package scan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ScanRouter struct{}

// InitScanRouter 初始化 扫码对比 路由信息
func (s *ScanRouter) InitScanRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	scan_compareRouter := Router.Group("scan_compare").Use(middleware.OperationRecord())
	scan_compareRouterWithoutRecord := Router.Group("scan_compare")
	scan_compareRouterWithoutAuth := PublicRouter.Group("scan_compare")
	{
		scan_compareRouter.POST("createScan", scan_compareApi.CreateScan)             // 新建扫码对比
		scan_compareRouter.DELETE("deleteScan", scan_compareApi.DeleteScan)           // 删除扫码对比
		scan_compareRouter.DELETE("deleteScanByIds", scan_compareApi.DeleteScanByIds) // 批量删除扫码对比
		scan_compareRouter.PUT("updateScan", scan_compareApi.UpdateScan)              // 更新扫码对比
	}
	{
		scan_compareRouterWithoutRecord.GET("findScan", scan_compareApi.FindScan)       // 根据ID获取扫码对比
		scan_compareRouterWithoutRecord.GET("getScanList", scan_compareApi.GetScanList) // 获取扫码对比列表
	}
	{
		scan_compareRouterWithoutAuth.GET("getScanPublic", scan_compareApi.GetScanPublic) // 扫码对比开放接口
		scan_compareRouterWithoutAuth.GET("", scan_compareApi.GetScanInfoPublic)          // 扫码对比开放接口
		scan_compareRouterWithoutAuth.POST("", scan_compareApi.PostScanInfoPublic)        // 扫码对比开放接口
	}
}
