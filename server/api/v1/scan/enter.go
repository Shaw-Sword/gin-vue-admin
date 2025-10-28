package scan

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ScanApi }

var scan_compareService = service.ServiceGroupApp.ScanServiceGroup.ScanService
