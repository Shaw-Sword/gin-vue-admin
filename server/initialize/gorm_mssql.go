package initialize

/*
 * @Author: 逆光飞翔 191180776@qq.com
 * @Date: 2022-12-08 17:25:49
 * @LastEditors: 逆光飞翔 191180776@qq.com
 * @LastEditTime: 2022-12-08 18:00:00
 * @FilePath: \server\initialize\gorm_mssql.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// GormMssql 初始化Mssql数据库
// Author [LouisZhang](191180776@qq.com)
func GormMssql() *gorm.DB {
	m := global.GVA_CONFIG.Mssql
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	global.GVA_LOG.Info("sqlserver正在尝试连接......")
	// 数据库配置
	general := m.GeneralDB
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(general)); err != nil {
		global.GVA_LOG.Sugar().Debugf("sqlserver连接失败:%v", err)
		return nil
	} else {
		global.GVA_LOG.Info("sqlserver连接成功")
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)

		// 数据库反向生成需要  正式环境移除
		// 为 SQL Server 2008 添加特殊配置
		//db.Set("gorm:table_options", "ENGINE=InnoDB")
		//
		//// 关键：禁用 OFFSET FETCH
		//sqlDB, err := db.DB()
		//if err == nil {
		//	// 设置连接参数，使用兼容 SQL Server 2008 的语法
		//	db.Exec("SET ANSI_NULLS ON")
		//	db.Exec("SET QUOTED_IDENTIFIER ON")
		//}

		return db
	}
}

// GormMssqlByConfig 初始化Mysql数据库用过传入配置
func GormMssqlByConfig(m config.Mssql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	// 数据库配置
	general := m.GeneralDB
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(general)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
