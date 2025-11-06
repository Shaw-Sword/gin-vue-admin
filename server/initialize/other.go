package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)

	cacheSaveDuration, err := utils.ParseDuration(global.GVA_CONFIG.ScanCompare.ExpiresTime)
	if err != nil {
		panic(err)
	}

	//      2025-10-29  扫码信息缓存
	global.ScanCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(cacheSaveDuration),
	)

	//file, err := os.Open("go.mod")
	//if err == nil && global.GVA_CONFIG.AutoCode.Module == "" {
	//	scanner := bufio.NewScanner(file)
	//	scanner.Scan()
	//	global.GVA_CONFIG.AutoCode.Module = strings.TrimPrefix(scanner.Text(), "module ")
	//}
}
