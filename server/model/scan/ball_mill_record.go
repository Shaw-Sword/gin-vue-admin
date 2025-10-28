package scan

import (
	"encoding/json"
	"time"
)

type BallMillRecordModel struct {
	Classification          string     `gorm:"column:classification" json:"classification"`
	TaskStatus              int        `gorm:"column:task_status;default:1" json:"task_status"`
	WeightFinishTime        *time.Time `gorm:"column:weight_finish_time" json:"weight_finish_time"`
	TaskID                  string     `gorm:"column:task_id;primaryKey" json:"task_id"`
	BaseNo                  string     `gorm:"column:base_no" json:"base_no"`
	BallMillNumber          string     `gorm:"column:ball_mill_number" json:"ball_mill_number"`
	CreateTime              *time.Time `gorm:"column:create_time" json:"create_time"`
	ExcelServerRCID         string     `gorm:"column:ExcelServerRCID" json:"-"`
	ExcelServerRC1          string     `gorm:"column:ExcelServerRC1" json:"-"`
	ExcelServerWIID         string     `gorm:"column:ExcelServerWIID" json:"-"`
	ExcelServerRTID         string     `gorm:"column:ExcelServerRTID" json:"-"`
	ExcelServerRN           int        `gorm:"column:ExcelServerRN" json:"-"`
	ExcelServerCN           int        `gorm:"column:ExcelServerCN" json:"-"`
	ExcelServerCHG          int        `gorm:"column:ExcelServerCHG" json:"-"`
	BarrelPlan1             float64    `gorm:"column:barrel_plan_1;not null" json:"barrel_plan_1"`
	BarrelActual1           float64    `gorm:"column:barrel_actual_1" json:"barrel_actual_1"`
	BarrelPlan2             float64    `gorm:"column:barrel_plan_2" json:"barrel_plan_2"`
	BarrelActual2           float64    `gorm:"column:barrel_actual_2" json:"barrel_actual_2"`
	BarrelPlan3             float64    `gorm:"column:barrel_plan_3" json:"barrel_plan_3"`
	BarrelActual3           float64    `gorm:"column:barrel_actual_3" json:"barrel_actual_3"`
	BarrelPlan4             float64    `gorm:"column:barrel_plan_4" json:"barrel_plan_4"`
	BarrelActual4           float64    `gorm:"column:barrel_actual_4" json:"barrel_actual_4"`
	BarrelPlan5             float64    `gorm:"column:barrel_plan_5" json:"barrel_plan_5"`
	BarrelActual5           float64    `gorm:"column:barrel_actual_5" json:"barrel_actual_5"`
	BarrelPlan6             float64    `gorm:"column:barrel_plan_6" json:"barrel_plan_6"`
	BarrelActual6           float64    `gorm:"column:barrel_actual_6" json:"barrel_actual_6"`
	BarrelPlan7             float64    `gorm:"column:barrel_plan_7" json:"barrel_plan_7"`
	BarrelActual7           float64    `gorm:"column:barrel_actual_7" json:"barrel_actual_7"`
	BarrelPlan8             float64    `gorm:"column:barrel_plan_8" json:"barrel_plan_8"`
	BarrelActual8           float64    `gorm:"column:barrel_actual_8" json:"barrel_actual_8"`
	BarrelPlan9             float64    `gorm:"column:barrel_plan_9" json:"barrel_plan_9"`
	BarrelActual9           float64    `gorm:"column:barrel_actual_9" json:"barrel_actual_9"`
	TotalWeight             float64    `gorm:"column:total_weight" json:"total_weight"`
	FeedFinishTime          *time.Time `gorm:"column:feed_finish_time" json:"feed_finish_time" format:"2006-01-02 15:04:05"`
	料1                      string     `gorm:"column:料1" json:"-"`
	料2                      string     `gorm:"column:料2" json:"-"`
	料3                      string     `gorm:"column:料3" json:"-"`
	料4                      string     `gorm:"column:料4" json:"-"`
	料5                      string     `gorm:"column:料5" json:"-"`
	料6                      string     `gorm:"column:料6" json:"-"`
	料7                      string     `gorm:"column:料7" json:"-"`
	料8                      string     `gorm:"column:料8" json:"-"`
	料9                      string     `gorm:"column:料9" json:"-"`
	ScanCodeCorrectNumber   int        `gorm:"column:scan_code_correct_number" json:"scan_code_correct_number"`
	AccumulateWeight        float64    `gorm:"column:accumulate_weight" json:"accumulate_weight"`
	ScanCodeErrorNumber     int        `gorm:"column:scan_code_error_number" json:"scan_code_error_number"`
	Ratio                   string     `gorm:"column:ratio" json:"ratio"`
	审核单                     string     `gorm:"column:审核单" json:"-"`
	PerPackageWeight        float64    `gorm:"column:per_package_weight" json:"per_package_weight"`
	TotalBarrelActualWeight float64    `gorm:"column:total_barrel_actual_weight" json:"total_barrel_actual_weight"`
	PriorityLevel           *int       `gorm:"column:priority_level" json:"priority_level"`
}

// MarshalJSON 自定义 JSON 输出
func (t BallMillRecordModel) MarshalJSON() ([]byte, error) {
	type Alias BallMillRecordModel
	return json.Marshal(&struct {
		CreateTime       *string `json:"create_time,omitempty"`
		WeightFinishTime *string `json:"weight_finish_time,omitempty"`
		FeedFinishTime   *string `json:"feed_finish_time,omitempty"`
		*Alias
	}{
		CreateTime:       formatTime(t.CreateTime),
		WeightFinishTime: formatTime(t.WeightFinishTime),
		FeedFinishTime:   formatTime(t.FeedFinishTime),
		Alias:            (*Alias)(&t),
	})
}

// TableName 重写表名
func (BallMillRecordModel) TableName() string {
	return "球磨报工记录单_主表"
}
