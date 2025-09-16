// 自动生成模板News
package huasu

import (
	"time"

	"github.com/kuihuar/huasu-admin/server/global"
)

// 新闻版块 结构体  News
type News struct {
	global.GVA_MODEL
	Title       *string    `json:"title" form:"title" gorm:"column:title;" binding:"required"`                   //新闻标题
	Category    *string    `json:"category" form:"category" gorm:"column:category;"`                             //新闻分类
	Summary     *string    `json:"summary" form:"summary" gorm:"column:summary;" binding:"required"`             //新闻摘要
	Content     *string    `json:"content" form:"content" gorm:"column:content;type:text;"`                      //新闻正文
	Cover       string     `json:"cover" form:"cover" gorm:"column:cover;"`                                      //封面图
	Source      *string    `json:"source" form:"source" gorm:"column:source;"`                                   //新闻来源
	Views       *int64     `json:"views" form:"views" gorm:"column:views;"`                                      //阅读量
	Publishtime *time.Time `json:"publishtime" form:"publishtime" gorm:"column:publishtime;" binding:"required"` //publishtime
	Status      *string    `json:"status" form:"status" gorm:"column:status;"`                                   //新闻状态
	Author      *string    `json:"author" form:"author" gorm:"column:author;"`                                   //作者
	Like        *int64     `json:"like" form:"like" gorm:"column:like;"`                                         //点赞数
	Dislike     *int64     `json:"dislike" form:"dislike" gorm:"column:dislike;"`                                //不点赞数
	Link        *string    `json:"link" form:"link" gorm:"column:link;"`                                         //链接
	Sort        *int64     `json:"sort" form:"sort" gorm:"default:1;column:sort;"`                               //排名
}

// TableName 新闻版块 News自定义表名 news
func (News) TableName() string {
	return "hs_news"
}
