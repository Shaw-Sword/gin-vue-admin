// 自动生成模板Scan
package scan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 扫码对比 结构体  Scan
type Scan struct {
	global.GVA_MODEL
	Recipe_content          *string `json:"recipe_content" form:"recipe_content" gorm:"comment:配方二维码内容;column:recipe_content;"`            //配方二维码内容
	Ball_mill_content       *string `json:"ball_mill_content" form:"ball_mill_content" gorm:"column:ball_mill_content;"`                   //球磨机二维码内容
	Compare_result          *string `json:"compare_result" form:"compare_result" gorm:"column:compare_result;"`                            //对比结果
	Ball_mill_report_record *string `json:"ball_mill_report_record" form:"ball_mill_report_record" gorm:"column:ball_mill_report_record;"` //球磨机报告记录单
	Weight_task_record      *string `json:"weight_task_record" form:"weight_task_record" gorm:"column:weight_task_record;"`                //称重任务单数据
}

// TableName 扫码对比 Scan自定义表名 tb_scan
func (Scan) TableName() string {
	return "tb_scan"
}
