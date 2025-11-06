package scan

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/scan"
	scanReq "github.com/flipped-aurora/gin-vue-admin/server/model/scan/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

const (
	DbCacheKey   = "db_result"   // 扫码缓存 数据库部分键名
	BallCacheKey = "ball_result" // 扫码缓存 球磨机 部分键名
)

type CacheDbInfo struct {
	TaskInfo      string
	RecordInfo    string
	Code          string
	TotalWeight   float64 // 任务重量
	CurrentWeight float64 // 当前称重重量
	TotalCount    int     // 总袋数
	CurrentCount  int     // 当前袋数
}

type ScanService struct {
	mu sync.Mutex
}

// CreateScan 创建扫码对比记录
// Author [yourname](https://github.com/yourname)
func (scan_compareService *ScanService) CreateScan(ctx context.Context, scan_compare *scan.Scan) (err error) {
	err = global.GVA_DB.Create(scan_compare).Error
	return err
}

// DeleteScan 删除扫码对比记录
// Author [yourname](https://github.com/yourname)
func (scan_compareService *ScanService) DeleteScan(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&scan.Scan{}, "id = ?", ID).Error
	return err
}

// DeleteScanByIds 批量删除扫码对比记录
// Author [yourname](https://github.com/yourname)
func (scan_compareService *ScanService) DeleteScanByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]scan.Scan{}, "id in ?", IDs).Error
	return err
}

// UpdateScan 更新扫码对比记录
// Author [yourname](https://github.com/yourname)
func (scan_compareService *ScanService) UpdateScan(ctx context.Context, scan_compare scan.Scan) (err error) {
	err = global.GVA_DB.Model(&scan.Scan{}).Where("id = ?", scan_compare.ID).Updates(&scan_compare).Error
	return err
}

// GetScan 根据ID获取扫码对比记录
// Author [yourname](https://github.com/yourname)
func (scan_compareService *ScanService) GetScan(ctx context.Context, ID string) (scan_compare scan.Scan, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&scan_compare).Error
	return
}

// GetScanInfoList 分页获取扫码对比记录
// Author [yourname](https://github.com/yourname)
func (scan_compareService *ScanService) GetScanInfoList(ctx context.Context, info scanReq.ScanSearch) (list []scan.Scan, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&scan.Scan{})
	var scan_compares []scan.Scan
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Recipe_content != nil && *info.Recipe_content != "" {
		db = db.Where("recipe_content LIKE ?", "%"+*info.Recipe_content+"%")
	}
	if info.Ball_mill_content != nil && *info.Ball_mill_content != "" {
		db = db.Where("ball_mill_content LIKE ?", "%"+*info.Ball_mill_content+"%")
	}
	if info.Compare_result != nil && *info.Compare_result != "" {
		db = db.Where("compare_result = ?", *info.Compare_result)
	}
	if info.Ball_mill_report_record != nil && *info.Ball_mill_report_record != "" {
		db = db.Where("ball_mill_report_record LIKE ?", "%"+*info.Ball_mill_report_record+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&scan_compares).Error
	return scan_compares, total, err
}
func (scan_compareService *ScanService) GetScanPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetCacheSaveDuration 获取扫码信息缓存时间
func GetCacheSaveDuration() time.Duration {
	cacheSaveDuration, err := utils.ParseDuration(global.GVA_CONFIG.ScanCompare.ExpiresTime)
	if err != nil {
		panic(err)
	}
	return cacheSaveDuration
}

// EIOCmdType 定义指令类型
type EIOCmdType int

const (
	CmdOpenGreen  EIOCmdType = iota + 1 // 1
	CmdCloseGreen                       // 2
	CmdOpenRed                          // 3
	CmdCloseRed                         // 4

)

// ToBytes 将指令转换为要发送的字节（这里用简单文本协议，也可用二进制/JSON）
func (c EIOCmdType) ToBytes() []byte {
	switch c {
	case CmdOpenGreen:
		return []byte("OPEN5\n")
	case CmdCloseGreen:
		return []byte("CLOSE5\n")
	case CmdOpenRed:
		return []byte("OPEN6\n")
	case CmdCloseRed:
		return []byte("CLOSE6\n")

	default:
		return []byte("CMD_UNKNOWN\n")
	}
}

// SendTcpCmd 发送 TCP 指令（一次性：连接 → 发送 → 关闭）
func SendTcpCmd(cmd EIOCmdType) {
	// 建立连接（带超时，避免卡死）
	addr := global.GVA_CONFIG.ScanCompare.EioIp + ":" + strconv.Itoa(global.GVA_CONFIG.ScanCompare.EioPort)
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		global.GVA_LOG.Sugar().Errorf("TCP连接EIO失败: %v", err.Error())
		//return fmt.Errorf("TCP连接EIO失败: %w", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			global.GVA_LOG.Sugar().Errorf("关闭TCP连接失败: %v", err.Error())
		}
	}(conn)
	// 设置写超时（可选，但推荐）
	err = conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		global.GVA_LOG.Sugar().Errorf("")
		return
	}
	// 发送指令
	_, err = conn.Write(cmd.ToBytes())
	if err != nil {
		global.GVA_LOG.Sugar().Errorf("TCP发送指令给EIO失败: %v", err.Error())
		//return fmt.Errorf("TCP发送指令给EIO失败: %w", err)
		return
	}
	// 注意：这里不读取响应（按你要求“不接受也行”）
	// 如果未来需要简单确认，可加 conn.Read(...)，但现在跳过
	global.GVA_LOG.Sugar().Infof("TCP发送指令到EIO设备成功")
	return
}

// 更新成功匹配的数据库信息
func updateOkTaskInfoByCode(info CacheDbInfo) {
	curTime := time.Now()
	if info.CurrentCount == info.TotalCount {
		db := global.E_MSSQL.Model(&scan.WeightTaskModel{TaskID: info.Code}).
			Select("AccumulateWeight", "ScanCodeCorrectNumber", "TaskStatus", "FeedFinishTime").
			Updates(map[string]interface{}{
				"AccumulateWeight":      gorm.Expr("ISNULL(accumulate_weight, 0) + ?", info.CurrentWeight),
				"ScanCodeCorrectNumber": gorm.Expr("ISNULL(scan_code_correct_number, 0) + 1"),
				"TaskStatus":            4,
				"FeedFinishTime":        &curTime,
			})
		if db.Error != nil {
			global.GVA_LOG.Sugar().Errorf("记录称重完成，更新数据库错误,%v", db.Error)
		}
		global.GVA_LOG.Info("记录称重信息完成，更新数据成功")

	} else {
		db := global.E_MSSQL.Model(&scan.WeightTaskModel{TaskID: info.Code}).
			Select("AccumulateWeight", "ScanCodeCorrectNumber").
			Updates(map[string]interface{}{
				"AccumulateWeight":      gorm.Expr("ISNULL(accumulate_weight, 0) + ?", info.CurrentWeight),
				"ScanCodeCorrectNumber": gorm.Expr("ISNULL(scan_code_correct_number, 0) + 1"),
			})
		if db.Error != nil {
			global.GVA_LOG.Sugar().Errorf("累积记录称重成功，更新数据库错误,%v", db.Error)
		}
		global.GVA_LOG.Info("累积记录称重成功，更新数据成功")

	}

}

// 更新错误匹配的数据库信息
func updateErrorTaskInfoByCode(info CacheDbInfo) {
	db := global.E_MSSQL.Model(&scan.WeightTaskModel{TaskID: info.Code}).
		Select("ScanCodeErrorNumber").
		Updates(map[string]interface{}{
			"ScanCodeErrorNumber": gorm.Expr("ISNULL(scan_code_error_number, 0) + 1"),
		})
	if db.Error != nil {
		global.GVA_LOG.Sugar().Errorf("累积记录称重扫码匹配失败信息，更新数据库错误,%v", db.Error)
	}
	global.GVA_LOG.Info("累积记录称重扫码匹配失败信息，记录成功")

}

// HandleScanInfoPublic 扫码后处理业务
func (s *ScanService) HandleScanInfoPublic(ctx context.Context, code string) error {
	s.mu.Lock()         // 上锁
	defer s.mu.Unlock() // 函数退出自动解锁

	// 💚 判断二维码类型   如果是球磨机二维码 ,查询缓存后,返回           否则
	if len(code) < 40 {
		// 查任务缓存是否存在
		dbCachedValue, found := global.ScanCache.Get(DbCacheKey)
		if found {
			// 类型断言
			dbInfo, ok := dbCachedValue.(CacheDbInfo)
			if !ok {
				fmt.Println("缓存类型错误")
				return fmt.Errorf("缓存类型错误")
			}
			if dbInfo.TaskInfo == code && dbInfo.RecordInfo == code { // 匹配成功
				global.GVA_LOG.Sugar().Infof("✔ 匹配成功,球磨机二维码信息：%s ,数据库基地等信息：%v", code, dbInfo)
				// 后续操作？？？
				updateOkTaskInfoByCode(dbInfo)

				global.ScanCache.Delete(DbCacheKey)
				global.ScanCache.Delete(BallCacheKey)

				// io模块 一致绿灯
				go func() {
					SendTcpCmd(CmdCloseRed)
					SendTcpCmd(CmdOpenGreen)
					time.Sleep(time.Second * 5)
					SendTcpCmd(CmdCloseGreen)
				}()

				return nil
			} else {
				global.GVA_LOG.Sugar().Errorf("❌匹配失败,称重任务信息：%s,球磨报告记录信息：%s,球磨机二维码信息：%s", dbInfo.TaskInfo, dbInfo.RecordInfo, code)

				updateErrorTaskInfoByCode(dbInfo)
				global.ScanCache.Delete(DbCacheKey)
				global.ScanCache.Delete(BallCacheKey)
				// io模块 不一致红灯
				SendTcpCmd(CmdCloseGreen)
				SendTcpCmd(CmdOpenRed)
				return fmt.Errorf("❌匹配失败,称重任务信息：%s,球磨报告记录信息：%s,球磨机二维码信息：%s", dbInfo.TaskInfo, dbInfo.RecordInfo, code)
			}
		} else { // 没有存 数据库查询到的数据  说明是第一次扫到球磨机二维码，缓存
			global.ScanCache.Set(BallCacheKey, code, GetCacheSaveDuration())
			global.GVA_LOG.Sugar().Infof("成功记录球磨机二维码信息 %s，等待匹配....", code)
			return nil
		}
	}

	// 长度大于40，说明是需要查询数据的二维码信息，需要截取 💚 查询 称重任务
	// 案例  2025/10/22 14:03_QMBG2510-00808_2025_Y3932_E1+LF1+LC1_19.9/19.88_1/1
	// 记录关键信息
	dbInfo := CacheDbInfo{}

	parts := strings.Split(code, "_") // 按下划线分割
	if len(parts) < 7 {
		global.GVA_LOG.Sugar().Errorf("格式错误,获取的二维码信息：%s", code)
		return fmt.Errorf("格式错误,获取的二维码信息：%s", code)

	}
	dbInfo.Code = parts[1]        // 第一个和第二个下划线之间的内容
	ratio1 := parts[len(parts)-2] // "19.9/19.88"
	ratio2 := parts[len(parts)-1] // "1/1"
	// 2. 分别按 '/' 拆分
	v1 := strings.Split(ratio1, "/")
	v2 := strings.Split(ratio2, "/")
	if len(v1) != 2 || len(v2) != 2 {
		global.GVA_LOG.Sugar().Errorf("重量和袋数格式错误,获取的二维码信息：%s", code)
		return fmt.Errorf("重量和袋数格式错误,获取的二维码信息：%s", code)
	}
	dbInfo.TotalWeight, _ = strconv.ParseFloat(v1[0], 64) //任务重量和称量重量_31.3/31.32
	dbInfo.CurrentWeight, _ = strconv.ParseFloat(v1[1], 64)
	dbInfo.TotalCount, _ = strconv.Atoi(v2[0])
	dbInfo.CurrentCount, _ = strconv.Atoi(v2[1]) // 总袋数和当前袋数_1/1

	var task scan.WeightTaskModel
	if err := global.E_MSSQL.Raw(`SELECT TOP 1 * FROM [dbo].[备料配方称重任务单_主表] WHERE task_id = ?`, dbInfo.Code).Scan(&task).Error; err != nil {
		global.GVA_LOG.Sugar().Errorf("查询称重任务失败,单号：%s", dbInfo.Code)
		return fmt.Errorf("查询单号：%s 称重任务失败: %w", dbInfo.Code, err)
	}
	if task.TaskID == "" {
		global.GVA_LOG.Sugar().Errorf("没有查询到称重任务,单号：%s", dbInfo.Code)
		return fmt.Errorf("查询单号：%s 称重任务为空", dbInfo.Code)
	}

	global.GVA_LOG.Sugar().Infof("查询成功,数据: %v", task)

	// 💚 查询 球磨报工记录
	var record scan.BallMillRecordModel
	if err := global.E_MSSQL.Raw(`SELECT TOP 1 * FROM [dbo].[球磨报工记录单_主表] WHERE 本单编码 = ?`, dbInfo.Code).Scan(&record).Error; err != nil {
		global.GVA_LOG.Sugar().Errorf("🔴查询球磨报工记录失败,单号：%s", dbInfo.Code)
		return fmt.Errorf("🔴查询单号：%s 球磨报工记录失败: %w", dbInfo.Code, err)
	}
	if record.Code == "" {
		global.GVA_LOG.Sugar().Errorf("🔴没有查询到球磨报工记录,单号：%s", dbInfo.Code)
		return fmt.Errorf("🔴查询单号：%s 球磨报工记录为空", dbInfo.Code)
	}
	global.GVA_LOG.Sugar().Infof("查询球磨报工记录成功,数据: %v", record)

	dbInfo.TaskInfo = string(task.BaseNo[len(task.BaseNo)-1]) + "0" + task.BallMillNumber[:2]
	dbInfo.RecordInfo = string(record.FactoryCode[len(record.FactoryCode)-1]) + "0" + record.BallMill[:2]

	// 存储二者的基地+球磨机号   一致的话再和 球磨机二维码缓存对比
	if ballCachedValue, found := global.ScanCache.Get(BallCacheKey); found {
		if ballCachedValue == dbInfo.TaskInfo && ballCachedValue == dbInfo.RecordInfo {
			// 匹配成功
			global.GVA_LOG.Sugar().Infof("✔ 匹配成功,单号信息：%s ,球磨机二维码信息：%s", dbInfo.Code, ballCachedValue)

			// 累积重量和成功次数
			updateOkTaskInfoByCode(dbInfo)

			global.ScanCache.Delete(DbCacheKey)
			global.ScanCache.Delete(BallCacheKey)
			// io模块 一致绿灯
			go func() {
				SendTcpCmd(CmdCloseRed)
				SendTcpCmd(CmdOpenGreen)
				time.Sleep(time.Second * 5)
				SendTcpCmd(CmdCloseGreen)
			}()

			return nil
		} else {
			global.GVA_LOG.Sugar().Errorf("❌匹配失败,称重任务信息：%s,球磨报告记录信息：%s,球磨机二维码信息：%s", dbInfo.TaskInfo, dbInfo.RecordInfo, ballCachedValue)

			// 累计失败次数
			updateErrorTaskInfoByCode(dbInfo)
			global.ScanCache.Delete(DbCacheKey)
			global.ScanCache.Delete(BallCacheKey)

			// io模块 不一致红灯
			SendTcpCmd(CmdCloseGreen)
			SendTcpCmd(CmdOpenRed)
			return fmt.Errorf("❌匹配失败,称重任务信息：%s,球磨报告记录信息：%s,球磨机二维码信息：%s", dbInfo.TaskInfo, dbInfo.RecordInfo, ballCachedValue)
		}

	} else { // 没有存过 球磨机二维码，说明是需要等待匹配
		global.ScanCache.Set(DbCacheKey, dbInfo, GetCacheSaveDuration())
		global.GVA_LOG.Sugar().Infof("成功记录数据库基地编号等信息 %v，等待球磨机二维码匹配....", dbInfo)
		return nil
	}

}
