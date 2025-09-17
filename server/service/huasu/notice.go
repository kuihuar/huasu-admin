package huasu

import (
	"context"

	"github.com/kuihuar/huasu-admin/server/global"
	"github.com/kuihuar/huasu-admin/server/model/huasu"
	huasuReq "github.com/kuihuar/huasu-admin/server/model/huasu/request"
)

type NoticeService struct{}

// CreateNotice 创建公告管理记录
// Author [yourname](https://github.com/yourname)
func (noticeService *NoticeService) CreateNotice(ctx context.Context, notice *huasu.Notice) (err error) {
	err = global.GVA_DB.Create(notice).Error
	return err
}

// DeleteNotice 删除公告管理记录
// Author [yourname](https://github.com/yourname)
func (noticeService *NoticeService) DeleteNotice(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&huasu.Notice{}, "id = ?", ID).Error
	return err
}

// DeleteNoticeByIds 批量删除公告管理记录
// Author [yourname](https://github.com/yourname)
func (noticeService *NoticeService) DeleteNoticeByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]huasu.Notice{}, "id in ?", IDs).Error
	return err
}

// UpdateNotice 更新公告管理记录
// Author [yourname](https://github.com/yourname)
func (noticeService *NoticeService) UpdateNotice(ctx context.Context, notice huasu.Notice) (err error) {
	err = global.GVA_DB.Model(&huasu.Notice{}).Where("id = ?", notice.ID).Updates(&notice).Error
	return err
}

// GetNotice 根据ID获取公告管理记录
// Author [yourname](https://github.com/yourname)
func (noticeService *NoticeService) GetNotice(ctx context.Context, ID string) (notice huasu.Notice, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&notice).Error
	return
}

// GetNoticeInfoList 分页获取公告管理记录
// Author [yourname](https://github.com/yourname)
func (noticeService *NoticeService) GetNoticeInfoList(ctx context.Context, info huasuReq.NoticeSearch) (list []huasu.Notice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&huasu.Notice{})
	var notices []huasu.Notice
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

	err = db.Find(&notices).Error
	return notices, total, err
}
func (noticeService *NoticeService) GetNoticePublic(ctx context.Context) ([]huasu.Notice, error) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
	db := global.GVA_DB.Model(&huasu.Notice{})
	var list []huasu.Notice
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("status= ?", 2).Order("top").Limit(3)

	err := db.Find(&list).Error
	return list, err
}

func (noticeService *NoticeService) GetNoticePublicList(ctx context.Context, info huasuReq.NoticeSearch) (list []huasu.Notice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&huasu.Notice{})
	var notices []huasu.Notice
	// 如果有条件搜索 下方会自动创建搜索语句
	// if len(info.CreatedAtRange) == 2 {
	// 	db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	// }

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&notices).Error
	return notices, total, err
}
