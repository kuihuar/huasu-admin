// 自动生成模板Notice
package huasu

import (
	"time"

	"github.com/kuihuar/huasu-admin/server/global"
)

// 公告管理 结构体  Notice
type Notice struct {
	global.GVA_MODEL
	Title      *string    `json:"title" form:"title" gorm:"column:title;" binding:"required"`                 //公告标题（必填）
	Summary    *string    `json:"summary" form:"summary" gorm:"column:summary;"`                              //公告摘要（可选）
	Content    *string    `json:"content" form:"content" gorm:"column:content;type:text;" binding:"required"` //公告正文（必填）
	Cover      string     `json:"cover" form:"cover" gorm:"column:cover;"`                                    //封面图 URL（可选）
	Source     *string    `json:"source" form:"source" gorm:"column:source;"`                                 //公告来源（可选）
	Status     *int64     `json:"status" form:"status" gorm:"column:status;" binding:"required"`              //发布状态（必填），控制公告是否对外展示，常见值： - 0：草稿（仅创建者可见，未提交审核） - 1：待审核（提交给管理员，未发布） - 2：已发布（全员 / 指定用户可见） - 3：已下架（历史公告，不再展示）
	Type       *int64     `json:"type" form:"type" gorm:"column:type;"`                                       //公告类型（必填，用于分类筛选），如： - 按业务分：系统公告、活动公告、通知公告、预警公告 - 按部门分：全员公告、技术部公告、财务部公告
	Top        *int64     `json:"top" form:"top" gorm:"column:top;"`                                          //是否置顶（可选），1 = 置顶，0 = 不置顶，置顶公告在列表页优先排序（如紧急系统通知）。
	Start_time *time.Time `json:"start_time" form:"start_time" gorm:"column:start_time;"`                     //生效时间（可选，定时发布）
	End_time   *time.Time `json:"end_time" form:"end_time" gorm:"column:end_time;"`                           //失效时间（可选，自动下架）
}

// TableName 公告管理 Notice自定义表名 notice
func (Notice) TableName() string {
	return "hs_notice"
}
