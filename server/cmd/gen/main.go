package main

import (
	"fmt"
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	global.GVA_VP = core.Viper()
	initialize.OtherInit()
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)

	global.E_MSSQL = initialize.MssqlGorm()
	if global.E_MSSQL == nil {
		log.Fatal("无法连接 MSSQL 数据库，请检查配置")
	}
	fmt.Println("已连接 MSSQL 数据库")

	// ========== 新版 gen 配置 ==========
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./server/model/external_models",
		ModelPkgPath: "external_models",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// 关键：为 SQL Server 2008 创建专用连接，禁用分页
	sql2008DB := global.E_MSSQL.Session(&gorm.Session{
		QueryFields: true,
	})

	// 设置 SQL Server 2008 兼容的查询选项
	sql2008DB = sql2008DB.Set("gorm:query_option", "TOP 1")

	// 绑定数据库
	g.UseDB(sql2008DB)

	// ========== 表映射：表名 → 模型名 ==========
	type TableMap struct {
		TableName string
		ModelName string
	}

	tables := []TableMap{
		{"备料配方称重任务单_主表", "BeiLiaoChengZhongTaskMain"},
		{"球磨报工记录单_主表", "QiuMoBaoGongMain"},
		// 继续添加...
	}

	for _, t := range tables {
		// 生成模型，指定表名和模型名
		m := g.GenerateModelAs(t.TableName, t.ModelName,
			// 可选：自定义字段标签、类型等
			gen.FieldTrimPrefix(""), // 如有共同前缀可去除
		)

		// 应用基础 CRUD 方法
		g.ApplyBasic(m)
	}

	// 执行生成
	g.Execute()

	fmt.Println("模型已成功生成到 ./server/model/external_models")
}
