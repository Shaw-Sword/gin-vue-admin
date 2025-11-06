package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ScanSearch struct {
	CreatedAtRange          []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Recipe_content          *string     `json:"recipe_content" form:"recipe_content"`
	Ball_mill_content       *string     `json:"ball_mill_content" form:"ball_mill_content"`
	Compare_result          *string     `json:"compare_result" form:"compare_result"`
	Ball_mill_report_record *string     `json:"ball_mill_report_record" form:"ball_mill_report_record"`
	request.PageInfo
}
