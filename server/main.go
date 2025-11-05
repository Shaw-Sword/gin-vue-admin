package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/gin-gonic/gin"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

//        //go:embed dist
//    var staticFiles embed.FS

// 这部分 @Tag 设置用于排序, 需要排序的接口请按照下面的格式添加
// swag init 对 @Tag 只会从入口文件解析, 默认 main.go
// 也可通过 --generalInfo flag 指定其他文件
// @Tag.Name        Base
// @Tag.Name        SysUser
// @Tag.Description 用户

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.8.6
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	// 👇 在程序最开始强制设置为 release 模式
	gin.SetMode(gin.ReleaseMode)
	// 初始化系统
	initializeSystem()
	// ====== 2. 嵌入 dist（现在就在当前目录下） ======
	//distFS, err := fs.Sub(staticFiles, "dist")
	//if err != nil {
	//	global.GVA_LOG.Fatal("无法加载前端资源:", err)
	//}
	// 运行服务器
	core.RunServer()
}

// initializeSystem 初始化系统所有组件
// 提取为单独函数以便于系统重载时调用
func initializeSystem() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库

	global.E_MSSQL = initialize.MssqlGorm() // 2025-10-28  外部数据库

	// initialize.Timer()
	// initialize.DBList()
	initialize.SetupHandlers() // 注册全局函数
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
	}
}
