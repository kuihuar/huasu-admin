
// 自动生成模板ConstructionManage
package huasu
import (
	"github.com/kuihuar/huasu-admin/server/global"
	"time"
	"gorm.io/datatypes"
)

// 工程管理 结构体  ConstructionManage
type ConstructionManage struct {
    global.GVA_MODEL
  Title  *string `json:"title" form:"title" gorm:"column:title;size:191;"`  //工程名
  Description  *string `json:"description" form:"description" gorm:"column:description;size:191;"`  //工程描述
  Type  *string `json:"type" form:"type" gorm:"column:type;size:191;"`  //工程类型
  Location  *string `json:"location" form:"location" gorm:"column:location;size:191;"`  //项目位置
  Owner  *string `json:"owner" form:"owner" gorm:"column:owner;size:191;"`  //责任人
  Design  *string `json:"design" form:"design" gorm:"column:design;size:191;"`  //设计单位
  Contractor  *string `json:"contractor" form:"contractor" gorm:"column:contractor;size:191;"`  //建造单位
  Management  *string `json:"management" form:"management" gorm:"column:management;size:191;"`  //管理单位
  Projectmanager  *string `json:"projectmanager" form:"projectmanager" gorm:"column:projectmanager;size:191;"`  //projectmanager字段
  Scheduleddate  *time.Time `json:"scheduleddate" form:"scheduleddate" gorm:"column:scheduleddate;"`  //开工日期
  Actualstartdate  *time.Time `json:"actualstartdate" form:"actualstartdate" gorm:"column:actualstartdate;"`  //实际开开日期
  Plannedfinishdate  *time.Time `json:"plannedfinishdate" form:"plannedfinishdate" gorm:"column:plannedfinishdate;"`  //计划结束日期
  Contractnumber  *string `json:"contractnumber" form:"contractnumber" gorm:"column:contractnumber;size:191;"`  //contractnumber字段
  Initiationdate  *time.Time `json:"initiationdate" form:"initiationdate" gorm:"column:initiationdate;"`  //开工日期
  Images  datatypes.JSON `json:"images" form:"images" gorm:"column:images;" swaggertype:"array,object"`  //项目展示图片
  Pdfs  datatypes.JSON `json:"pdfs" form:"pdfs" gorm:"column:pdfs;" swaggertype:"array,object"`  //项目文档
  Video  string `json:"video" form:"video" gorm:"column:video;size:191;"`  //项目视频
  Vudgetedcost  *int64 `json:"vudgetedcost" form:"vudgetedcost" gorm:"column:vudgetedcost;"`  //vudgetedcost字段
  Actualcost  *int64 `json:"actualcost" form:"actualcost" gorm:"column:actualcost;"`  //实际花费
  Estimatecompletion  *int64 `json:"estimatecompletion" form:"estimatecompletion" gorm:"column:estimatecompletion;"`  //预估花费
  Costunit  *string `json:"costunit" form:"costunit" gorm:"column:costunit;size:191;"`  //花费单位
}


// TableName 工程管理 ConstructionManage自定义表名 hs_construction_manage
func (ConstructionManage) TableName() string {
    return "hs_construction_manage"
}





