// 自动生成模板CarouselManage
package huasu

import (
	"github.com/kuihuar/huasu-admin/server/global"
)

// 轮播图管理 结构体  CarouselManage
type CarouselManage struct {
	global.GVA_MODEL
	Title       *string `json:"title" form:"title" gorm:"comment:轮播图标题;column:title;" binding:"required"`             //轮播图标题
	Image       string  `json:"image" form:"image" gorm:"comment:轮播图片的存储路径或 CDN 链接;column:image;" binding:"required"` //轮播图片的存储路径或 CDN 链接
	Link        *string `json:"link" form:"link" gorm:"column:link;" binding:"required"`                              //点击轮播图后的跳转链接
	Sort        *int64  `json:"sort" form:"sort" gorm:"column:sort;" binding:"required"`                              //排序权重
	Status      *int64  `json:"status" form:"status" gorm:"column:status;" binding:"required"`                        //状态（如：1 - 启用、0 - 禁用、2 - 草稿，控制是否在前端显示）
	Business    *string `json:"business" form:"business" gorm:"index;column:business;" binding:"required"`            //关联业务模块
	Description *string `json:"description" form:"description" gorm:"column:description;"`                            //轮播图描述（管理后台备注，如 “618 活动主视觉”）
	Alt         *string `json:"alt" form:"alt" gorm:"column:alt;"`                                                    //图片的 alt 属性
	Target      *string `json:"target" form:"target" gorm:"default:_self;column:target;"`                             //链接打开方式
	CreatedBy   *int64  `json:"created_by" form:"created_by" gorm:"column:created_by;"`                               //创建人 ID
}

// TableName 轮播图管理 CarouselManage自定义表名 carousel
func (CarouselManage) TableName() string {
	return "hs_carousel"
}
