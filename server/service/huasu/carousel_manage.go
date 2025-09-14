package huasu

import (
	"context"

	"github.com/kuihuar/huasu-admin/server/global"
	"github.com/kuihuar/huasu-admin/server/model/huasu"
	huasuReq "github.com/kuihuar/huasu-admin/server/model/huasu/request"
)

type CarouselManageService struct{}

// CreateCarouselManage 创建轮播图管理记录
// Author [yourname](https://github.com/yourname)
func (CarouselService *CarouselManageService) CreateCarouselManage(ctx context.Context, Carousel *huasu.CarouselManage) (err error) {
	err = global.GVA_DB.Create(Carousel).Error
	return err
}

// DeleteCarouselManage 删除轮播图管理记录
// Author [yourname](https://github.com/yourname)
func (CarouselService *CarouselManageService) DeleteCarouselManage(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&huasu.CarouselManage{}, "id = ?", ID).Error
	return err
}

// DeleteCarouselManageByIds 批量删除轮播图管理记录
// Author [yourname](https://github.com/yourname)
func (CarouselService *CarouselManageService) DeleteCarouselManageByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]huasu.CarouselManage{}, "id in ?", IDs).Error
	return err
}

// UpdateCarouselManage 更新轮播图管理记录
// Author [yourname](https://github.com/yourname)
func (CarouselService *CarouselManageService) UpdateCarouselManage(ctx context.Context, Carousel huasu.CarouselManage) (err error) {
	err = global.GVA_DB.Model(&huasu.CarouselManage{}).Where("id = ?", Carousel.ID).Updates(&Carousel).Error
	return err
}

// GetCarouselManage 根据ID获取轮播图管理记录
// Author [yourname](https://github.com/yourname)
func (CarouselService *CarouselManageService) GetCarouselManage(ctx context.Context, ID string) (Carousel huasu.CarouselManage, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&Carousel).Error
	return
}

// GetCarouselManageInfoList 分页获取轮播图管理记录
// Author [yourname](https://github.com/yourname)
func (CarouselService *CarouselManageService) GetCarouselManageInfoList(ctx context.Context, info huasuReq.CarouselManageSearch) (list []huasu.CarouselManage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&huasu.CarouselManage{})
	var Carousels []huasu.CarouselManage
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["sort"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Carousels).Error
	return Carousels, total, err
}
func (CarouselService *CarouselManageService) GetCarouselManagePublic(ctx context.Context) (list []huasu.CarouselManage, err error) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
	db := global.GVA_DB.Model(&huasu.CarouselManage{})
	var Carousels []huasu.CarouselManage
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("business= ?", "home").Order("sort").Limit(4)

	err = db.Find(&Carousels).Error
	return Carousels, err
}
