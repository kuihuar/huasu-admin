package huasu

import (
	"context"

	"github.com/kuihuar/huasu-admin/server/global"
	"github.com/kuihuar/huasu-admin/server/model/huasu"
	huasuReq "github.com/kuihuar/huasu-admin/server/model/huasu/request"
)

type NewsService struct{}

// CreateNews 创建新闻版块记录
// Author [yourname](https://github.com/yourname)
func (newsService *NewsService) CreateNews(ctx context.Context, news *huasu.News) (err error) {
	err = global.GVA_DB.Create(news).Error
	return err
}

// DeleteNews 删除新闻版块记录
// Author [yourname](https://github.com/yourname)
func (newsService *NewsService) DeleteNews(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&huasu.News{}, "id = ?", ID).Error
	return err
}

// DeleteNewsByIds 批量删除新闻版块记录
// Author [yourname](https://github.com/yourname)
func (newsService *NewsService) DeleteNewsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]huasu.News{}, "id in ?", IDs).Error
	return err
}

// UpdateNews 更新新闻版块记录
// Author [yourname](https://github.com/yourname)
func (newsService *NewsService) UpdateNews(ctx context.Context, news huasu.News) (err error) {
	err = global.GVA_DB.Model(&huasu.News{}).Where("id = ?", news.ID).Updates(&news).Error
	return err
}

// GetNews 根据ID获取新闻版块记录
// Author [yourname](https://github.com/yourname)
func (newsService *NewsService) GetNews(ctx context.Context, ID string) (news huasu.News, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&news).Error
	return
}

// GetNewsInfoList 分页获取新闻版块记录
// Author [yourname](https://github.com/yourname)
func (newsService *NewsService) GetNewsInfoList(ctx context.Context, info huasuReq.NewsSearch) (list []huasu.News, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&huasu.News{})
	var newss []huasu.News
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&newss).Error
	return newss, total, err
}
func (newsService *NewsService) GetNewsPublic(ctx context.Context, info huasuReq.NewsSearch) (list []huasu.News, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&huasu.News{})
	var newss []huasu.News
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&newss).Error
	return newss, total, err
}
