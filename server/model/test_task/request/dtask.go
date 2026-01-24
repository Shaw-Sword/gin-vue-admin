
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BTaskSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Task_no  *string `json:"task_no" form:"task_no"` 
      Num  *int `json:"num" form:"num"` 
    request.PageInfo
}
