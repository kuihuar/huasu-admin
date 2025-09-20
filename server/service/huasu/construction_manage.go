package huasu

import (
	"context"

	"github.com/kuihuar/huasu-admin/server/global"
	"github.com/kuihuar/huasu-admin/server/model/huasu"
	huasuReq "github.com/kuihuar/huasu-admin/server/model/huasu/request"
)

type ConstructionManageService struct{}

// CreateConstructionManage 创建工程管理记录
// Author [yourname](https://github.com/yourname)
func (constructionManageService *ConstructionManageService) CreateConstructionManage(ctx context.Context, constructionManage *huasu.ConstructionManage) (err error) {
	err = global.GVA_DB.Create(constructionManage).Error
	return err
}

// DeleteConstructionManage 删除工程管理记录
// Author [yourname](https://github.com/yourname)
func (constructionManageService *ConstructionManageService) DeleteConstructionManage(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&huasu.ConstructionManage{}, "id = ?", ID).Error
	return err
}

// DeleteConstructionManageByIds 批量删除工程管理记录
// Author [yourname](https://github.com/yourname)
func (constructionManageService *ConstructionManageService) DeleteConstructionManageByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]huasu.ConstructionManage{}, "id in ?", IDs).Error
	return err
}

// UpdateConstructionManage 更新工程管理记录
// Author [yourname](https://github.com/yourname)
func (constructionManageService *ConstructionManageService) UpdateConstructionManage(ctx context.Context, constructionManage huasu.ConstructionManage) (err error) {
	err = global.GVA_DB.Model(&huasu.ConstructionManage{}).Where("id = ?", constructionManage.ID).Updates(&constructionManage).Error
	return err
}

// GetConstructionManage 根据ID获取工程管理记录
// Author [yourname](https://github.com/yourname)
func (constructionManageService *ConstructionManageService) GetConstructionManage(ctx context.Context, ID string) (constructionManage huasu.ConstructionManage, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&constructionManage).Error
	return
}

// GetConstructionManageInfoList 分页获取工程管理记录
// Author [yourname](https://github.com/yourname)
func (constructionManageService *ConstructionManageService) GetConstructionManageInfoList(ctx context.Context, info huasuReq.ConstructionManageSearch) (list []huasu.ConstructionManage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&huasu.ConstructionManage{})
	var constructionManages []huasu.ConstructionManage
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

	err = db.Find(&constructionManages).Error
	return constructionManages, total, err
}
func (constructionManageService *ConstructionManageService) GetConstructionManagePublic(ctx context.Context, search huasuReq.ConstructionManageSearch) (list []huasu.ConstructionManage, total int64, err error) {
	limit := search.PageSize
	offset := search.PageSize * (search.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&huasu.ConstructionManage{})
	var constructionManages []huasu.ConstructionManage
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(search.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", search.CreatedAtRange[0], search.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&constructionManages).Error
	return constructionManages, total, err
}
