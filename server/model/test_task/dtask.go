
// 自动生成模板BTask
package test_task
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 测试任务 结构体  BTask
type BTask struct {
    global.GVA_MODEL
  Task_no  *string `json:"task_no" form:"task_no" gorm:"column:task_no;"`  //任务编号
  Num  *int64 `json:"num" form:"num" gorm:"default:0;column:num;"`  //数量
  Weight  *float64 `json:"weight" form:"weight" gorm:"default:0;column:weight;"`  //重量
  Test_task_status  *string `json:"test_task_status" form:"test_task_status" gorm:"column:test_task_status;"`  //任务类型
}


// TableName 测试任务 BTask自定义表名 tb_jy_task
func (BTask) TableName() string {
    return "tb_jy_task"
}





