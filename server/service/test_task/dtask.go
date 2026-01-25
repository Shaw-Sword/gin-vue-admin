package test_task

import (
	"context"
	"gva/global"
	"gva/model/test_task"
	test_taskReq "gva/model/test_task/request"
)

type BTaskService struct{}

// CreateBTask 创建测试任务记录
// Author [yourname](https://github.com/yourname)
func (CTaskService *BTaskService) CreateBTask(ctx context.Context, CTask *test_task.BTask) (err error) {
	err = global.GVA_DB.Create(CTask).Error
	return err
}

// DeleteBTask 删除测试任务记录
// Author [yourname](https://github.com/yourname)
func (CTaskService *BTaskService) DeleteBTask(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&test_task.BTask{}, "id = ?", ID).Error
	return err
}

// DeleteBTaskByIds 批量删除测试任务记录
// Author [yourname](https://github.com/yourname)
func (CTaskService *BTaskService) DeleteBTaskByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]test_task.BTask{}, "id in ?", IDs).Error
	return err
}

// UpdateBTask 更新测试任务记录
// Author [yourname](https://github.com/yourname)
func (CTaskService *BTaskService) UpdateBTask(ctx context.Context, CTask test_task.BTask) (err error) {
	err = global.GVA_DB.Model(&test_task.BTask{}).Where("id = ?", CTask.ID).Updates(&CTask).Error
	return err
}

// GetBTask 根据ID获取测试任务记录
// Author [yourname](https://github.com/yourname)
func (CTaskService *BTaskService) GetBTask(ctx context.Context, ID string) (CTask test_task.BTask, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&CTask).Error
	return
}

// GetBTaskInfoList 分页获取测试任务记录
// Author [yourname](https://github.com/yourname)
func (CTaskService *BTaskService) GetBTaskInfoList(ctx context.Context, info test_taskReq.BTaskSearch) (list []test_task.BTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&test_task.BTask{})
	var CTasks []test_task.BTask
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Task_no != nil && *info.Task_no != "" {
		db = db.Where("task_no LIKE ?", "%"+*info.Task_no+"%")
	}
	if info.Num != nil {
		db = db.Where("num = ?", *info.Num)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&CTasks).Error
	return CTasks, total, err
}
func (CTaskService *BTaskService) GetBTaskPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// Test_fn_t1 测试方法1
// Author [yourname](https://github.com/yourname)
func (CTaskService *BTaskService) Test_fn_t1(ctx context.Context) (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&test_task.BTask{})
	return db.Error
}
