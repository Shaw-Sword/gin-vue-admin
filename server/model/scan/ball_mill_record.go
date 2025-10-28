package scan

import (
	"encoding/json"
	"time"
)

// BallMillRecordModel 球磨报工记录单主表
type BallMillRecordModel struct {
	FeedReportDate          *time.Time `gorm:"column:投料报工日期" json:"feed_report_date"`
	Code                    string     `gorm:"column:本单编码;size:50" json:"code"`
	Creator                 string     `gorm:"column:制单人;size:20" json:"creator"`
	CreateDate              *time.Time `gorm:"column:制单日期" json:"create_date"`
	ExcelServerRCID         string     `gorm:"column:ExcelServerRCID;size:20" json:"-"`
	ExcelServerRC1          string     `gorm:"column:ExcelServerRC1;size:20" json:"-"`
	ExcelServerWIID         string     `gorm:"column:ExcelServerWIID;size:20" json:"-"`
	ExcelServerRTID         string     `gorm:"column:ExcelServerRTID;size:20" json:"-"`
	ExcelServerRN           *int       `gorm:"column:ExcelServerRN" json:"-"`
	ExcelServerCN           *int       `gorm:"column:ExcelServerCN" json:"-"`
	ExcelServerCHG          *int       `gorm:"column:ExcelServerCHG" json:"-"`
	BallMill                string     `gorm:"column:球磨机;size:100" json:"ball_mill"`
	TaskNo                  string     `gorm:"column:任务单号;size:50" json:"task_no"`
	Grade                   string     `gorm:"column:牌号;size:100" json:"grade"`
	RawMaterialBatchA       string     `gorm:"column:原材料批号A;size:100" json:"raw_material_batch_a"`
	MaterialNo              string     `gorm:"column:料号;size:100" json:"material_no"`
	BaseRatioT              string     `gorm:"column:基料配比T;size:100" json:"base_ratio_t"`
	BaseRatioRawA           *float64   `gorm:"column:基料配比中原料A;precision:20;scale:3" json:"base_ratio_raw_a"`
	BaseRatioCrushedT       *float64   `gorm:"column:基料配比中破碎料T;precision:20;scale:3" json:"base_ratio_crushed_t"`
	CrushedWorkerNo         string     `gorm:"column:破碎料工工号;size:100" json:"crushed_worker_no"`
	BaseRatioMagneticT      *float64   `gorm:"column:基料配比中磁选料T;precision:20;scale:3" json:"base_ratio_magnetic_t"`
	MagneticWorkerNoT       string     `gorm:"column:磁选料工号T;size:100" json:"magnetic_worker_no_t"`
	StandardTime            *float64   `gorm:"column:标准时间;precision:20;scale:3" json:"standard_time"`
	FeedWeightT             *float64   `gorm:"column:投料重量T;precision:20;scale:3" json:"feed_weight_t"`
	FeedWorkerNo            string     `gorm:"column:投料工工号;size:100" json:"feed_worker_no"`
	PlanFormulaNo           string     `gorm:"column:计划配方编号;size:100" json:"plan_formula_no"`
	StartTime               *time.Time `gorm:"column:开机时间" json:"start_time"`
	Accumulated1Hour        *float64   `gorm:"column:累计1小时;precision:20;scale:3" json:"accumulated_1_hour"`
	SecondStartTime         *time.Time `gorm:"column:二次开机" json:"second_start_time"`
	Accumulated2Hour        *float64   `gorm:"column:累计2小时;precision:20;scale:3" json:"accumulated_2_hour"`
	ThirdStartTime          *time.Time `gorm:"column:三次开机" json:"third_start_time"`
	DischargeWeightT        *float64   `gorm:"column:出料重量T;precision:20;scale:3" json:"discharge_weight_t"`
	DischargeWorkerNo       string     `gorm:"column:出料工工号;size:100" json:"discharge_worker_no"`
	DischargeTowerNo        string     `gorm:"column:出料指定沉淀塔号;size:100" json:"discharge_tower_no"`
	DischargedTowerNo       string     `gorm:"column:脱料出的塔号;size:100" json:"discharged_tower_no"`
	DischargeBatchNo        string     `gorm:"column:出料后的批次号;size:100" json:"discharge_batch_no"`
	FeedShift               string     `gorm:"column:投料班次;size:20" json:"feed_shift"`
	RegrindReportNo         string     `gorm:"column:返磨报工单;size:50" json:"regrind_report_no"`
	ShutdownTime            *time.Time `gorm:"column:关机时间" json:"shutdown_time"`
	SecondShutdownTime      *time.Time `gorm:"column:二次关机" json:"second_shutdown_time"`
	ThirdShutdownTime       *time.Time `gorm:"column:三次关机" json:"third_shutdown_time"`
	Accumulated3Hour        *float64   `gorm:"column:累计3小时;precision:20;scale:3" json:"accumulated_3_hour"`
	TotalAccumulatedTime    *float64   `gorm:"column:总累计时间;precision:20;scale:3" json:"total_accumulated_time"`
	FactoryBuilding         string     `gorm:"column:厂房;size:20" json:"factory_building"`
	Status                  string     `gorm:"column:状态;size:20" json:"status"`
	ProcessStatus           string     `gorm:"column:处理状态;size:20" json:"process_status"`
	RegrindBatchNo          string     `gorm:"column:返磨批次号;size:100" json:"regrind_batch_no"`
	FeedPlanRemark          string     `gorm:"column:投料计划备注;size:100" json:"feed_plan_remark"`
	RawMaterialBatchB       string     `gorm:"column:原材料批号B;size:100" json:"raw_material_batch_b"`
	RawMaterialBatchC       string     `gorm:"column:原材料批号C;size:100" json:"raw_material_batch_c"`
	BaseRatioRawB           *float64   `gorm:"column:基料配比中原料B;precision:20;scale:3" json:"base_ratio_raw_b"`
	BaseRatioRawC           *float64   `gorm:"column:基料配比中原料C;precision:20;scale:3" json:"base_ratio_raw_c"`
	BaseRatioReturnT        *float64   `gorm:"column:基料配比中回料T;precision:20;scale:3" json:"base_ratio_return_t"`
	RegrindTowerNo          string     `gorm:"column:返磨沉淀塔号;size:100" json:"regrind_tower_no"`
	DischargeSpecifiedDate  *time.Time `gorm:"column:出料指定日期" json:"discharge_specified_date"`
	RawMaterialA            string     `gorm:"column:原材料A;size:100" json:"raw_material_a"`
	RawMaterialB            string     `gorm:"column:原材料B;size:100" json:"raw_material_b"`
	RawMaterialC            string     `gorm:"column:原材料C;size:100" json:"raw_material_c"`
	FeedReporter            string     `gorm:"column:投料报工人;size:20" json:"feed_reporter"`
	OriginalWeight          *float64   `gorm:"column:原重量;precision:20;scale:3" json:"original_weight"`
	DischargeRemark         string     `gorm:"column:出料备注;size:100" json:"discharge_remark"`
	FeedRemark              string     `gorm:"column:投料备注;size:100" json:"feed_remark"`
	FormulaNo               string     `gorm:"column:配方编号;size:100" json:"formula_no"`
	SourcePlanNo            string     `gorm:"column:来源计划单;size:50" json:"source_plan_no"`
	SerialNo                string     `gorm:"column:流水号;size:50" json:"serial_no"`
	FactoryCode             string     `gorm:"column:厂房代码;size:100" json:"factory_code"`
	FormulaProcessCode      string     `gorm:"column:配方工艺代码;size:100" json:"formula_process_code"`
	AutoJudgeResult         string     `gorm:"column:自动判定结果;size:100" json:"auto_judge_result"`
	InspectionRemark        string     `gorm:"column:检验备注;size:500" json:"inspection_remark"`
	BrMT                    *float64   `gorm:"column:BrMT;precision:20;scale:3" json:"br_mt"`
	HcbkAm                  *float64   `gorm:"column:HcbkAm;precision:20;scale:3" json:"hcbk_am"`
	HcjkAm                  *float64   `gorm:"column:HcjkAm;precision:20;scale:3" json:"hcjk_am"`
	BHm                     *float64   `gorm:"column:BHm;precision:20;scale:3" json:"b_hm"`
	ShrinkageRatio          *float64   `gorm:"column:收缩比;precision:20;scale:3" json:"shrinkage_ratio"`
	Temperature             *float64   `gorm:"column:温度;precision:20;scale:3" json:"temperature"`
	RecordNo                string     `gorm:"column:记录号;size:100" json:"record_no"`
	MaterialSinterKilnNo    string     `gorm:"column:料号烧结窑炉号;size:100" json:"material_sinter_kiln_no"`
	FeedPerformanceTestTime *time.Time `gorm:"column:投料性能测试报工时间" json:"feed_performance_test_time"`
	WorkCenterCode          string     `gorm:"column:工作中心编码;size:50" json:"work_center_code"`
	RawMaterialACode        string     `gorm:"column:原材料A物资编码;size:50" json:"raw_material_a_code"`
	RawMaterialBCode        string     `gorm:"column:原材料B物资编码;size:50" json:"raw_material_b_code"`
	RawMaterialCCode        string     `gorm:"column:原材料C物资编码;size:50" json:"raw_material_c_code"`
	CrushedBatchNo          string     `gorm:"column:破碎料批号;size:100" json:"crushed_batch_no"`
	MagneticBatchNo         string     `gorm:"column:磁选料批号;size:100" json:"magnetic_batch_no"`
	OtherBatchNo            string     `gorm:"column:其他料批号;size:100" json:"other_batch_no"`
	OperationGuideCode      string     `gorm:"column:作业指导书编码;size:50" json:"operation_guide_code"`
	CrushedMaterialCode     string     `gorm:"column:破碎料物资编码;size:50" json:"crushed_material_code"`
	MagneticMaterialCode    string     `gorm:"column:磁选料物资编码;size:50" json:"magnetic_material_code"`
	OtherMaterialCode       string     `gorm:"column:其他料物资编码;size:50" json:"other_material_code"`
	Size                    *float64   `gorm:"column:尺寸;precision:20;scale:3" json:"size"`
	MoldSize                *float64   `gorm:"column:模具尺寸;precision:20;scale:3" json:"mold_size"`
	Granularity             *float64   `gorm:"column:粒度;precision:20;scale:3" json:"granularity"`
	PHValue                 *float64   `gorm:"column:PH值;precision:20;scale:3" json:"ph_value"`
	PerformanceTester       string     `gorm:"column:性能测试人;size:100" json:"performance_tester"`
	PlanFeedDate            *time.Time `gorm:"column:计划投料日期" json:"plan_feed_date"`
}

// MarshalJSON 自定义 JSON 输出
func (t BallMillRecordModel) MarshalJSON() ([]byte, error) {
	type Alias BallMillRecordModel
	return json.Marshal(&struct {
		FeedReportDate          *string `json:"feed_report_date,omitempty"`
		CreateDate              *string `json:"create_date,omitempty"`
		ShutdownTime            *string `json:"shutdown_time,omitempty"`
		SecondShutdownTime      *string `json:"second_shutdown_time,omitempty"`
		ThirdShutdownTime       *string `json:"third_shutdown_time,omitempty"`
		StartTime               *string `json:"start_time,omitempty"`
		SecondStartTime         *string `json:"second_start_time,omitempty"`
		ThirdStartTime          *string `json:"third_start_time,omitempty"`
		DischargeSpecifiedDate  *string `json:"discharge_specified_date,omitempty"`
		FeedPerformanceTestTime *string `json:"feed_performance_test_time,omitempty"`
		PlanFeedDate            *string `json:"plan_feed_date,omitempty"`

		*Alias
	}{
		FeedReportDate:          formatTime(t.FeedReportDate),
		CreateDate:              formatTime(t.CreateDate),
		ShutdownTime:            formatTime(t.ShutdownTime),
		SecondShutdownTime:      formatTime(t.SecondShutdownTime),
		ThirdShutdownTime:       formatTime(t.ThirdShutdownTime),
		StartTime:               formatTime(t.StartTime),
		SecondStartTime:         formatTime(t.SecondStartTime),
		ThirdStartTime:          formatTime(t.ThirdStartTime),
		DischargeSpecifiedDate:  formatTime(t.DischargeSpecifiedDate),
		FeedPerformanceTestTime: formatTime(t.FeedPerformanceTestTime),
		PlanFeedDate:            formatTime(t.PlanFeedDate),
		Alias:                   (*Alias)(&t),
	})
}

// TableName 重写表名
func (BallMillRecordModel) TableName() string {
	return "球磨报工记录单_主表"
}
