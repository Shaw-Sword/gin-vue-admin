package scan

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/scan"
	scanReq "github.com/flipped-aurora/gin-vue-admin/server/model/scan/request"
)

type ScanService struct{}

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
