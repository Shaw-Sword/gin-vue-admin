package scan

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ScanRouter }

var scan_compareApi = api.ApiGroupApp.ScanApiGroup.ScanApi
