package test_task

import (
	"github.com/gin-gonic/gin"
	"gva/middleware"
)

type BTaskRouter struct{}

func (s *BTaskRouter) InitBTaskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	CTaskRouter := Router.Group("CTask").Use(middleware.OperationRecord())
	CTaskRouterWithoutRecord := Router.Group("CTask")
	CTaskRouterWithoutAuth := PublicRouter.Group("CTask")
	{
		CTaskRouter.POST("createBTask", CTaskApi.CreateBTask)
		CTaskRouter.DELETE("deleteBTask", CTaskApi.DeleteBTask)
		CTaskRouter.DELETE("deleteBTaskByIds", CTaskApi.DeleteBTaskByIds)
		CTaskRouter.PUT("updateBTask", CTaskApi.UpdateBTask)
		CTaskRouter.POST("test_fn1", CTaskApi.Test_fn_t1)
	}
	{
		CTaskRouterWithoutRecord.GET("findBTask", CTaskApi.FindBTask)
		CTaskRouterWithoutRecord.GET("getBTaskList", CTaskApi.GetBTaskList)
	}
	{
		CTaskRouterWithoutAuth.GET("getBTaskPublic", CTaskApi.GetBTaskPublic)
	}
}
