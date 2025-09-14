package huasu

import (
	"github.com/gin-gonic/gin"
	"github.com/kuihuar/huasu-admin/server/global"
	"github.com/kuihuar/huasu-admin/server/model/common/response"
	"github.com/kuihuar/huasu-admin/server/model/huasu"
	huasuReq "github.com/kuihuar/huasu-admin/server/model/huasu/request"
	"go.uber.org/zap"
)

type CarouselManageApi struct{}

// CreateCarouselManage 创建轮播图管理
// @Tags CarouselManage
// @Summary 创建轮播图管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body huasu.CarouselManage true "创建轮播图管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Carousel/createCarouselManage [post]
func (CarouselApi *CarouselManageApi) CreateCarouselManage(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var Carousel huasu.CarouselManage
	err := c.ShouldBindJSON(&Carousel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = CarouselService.CreateCarouselManage(ctx, &Carousel)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCarouselManage 删除轮播图管理
// @Tags CarouselManage
// @Summary 删除轮播图管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body huasu.CarouselManage true "删除轮播图管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Carousel/deleteCarouselManage [delete]
func (CarouselApi *CarouselManageApi) DeleteCarouselManage(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := CarouselService.DeleteCarouselManage(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCarouselManageByIds 批量删除轮播图管理
// @Tags CarouselManage
// @Summary 批量删除轮播图管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Carousel/deleteCarouselManageByIds [delete]
func (CarouselApi *CarouselManageApi) DeleteCarouselManageByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := CarouselService.DeleteCarouselManageByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCarouselManage 更新轮播图管理
// @Tags CarouselManage
// @Summary 更新轮播图管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body huasu.CarouselManage true "更新轮播图管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Carousel/updateCarouselManage [put]
func (CarouselApi *CarouselManageApi) UpdateCarouselManage(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var Carousel huasu.CarouselManage
	err := c.ShouldBindJSON(&Carousel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = CarouselService.UpdateCarouselManage(ctx, Carousel)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCarouselManage 用id查询轮播图管理
// @Tags CarouselManage
// @Summary 用id查询轮播图管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询轮播图管理"
// @Success 200 {object} response.Response{data=huasu.CarouselManage,msg=string} "查询成功"
// @Router /Carousel/findCarouselManage [get]
func (CarouselApi *CarouselManageApi) FindCarouselManage(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reCarousel, err := CarouselService.GetCarouselManage(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reCarousel, c)
}

// GetCarouselManageList 分页获取轮播图管理列表
// @Tags CarouselManage
// @Summary 分页获取轮播图管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query huasuReq.CarouselManageSearch true "分页获取轮播图管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Carousel/getCarouselManageList [get]
func (CarouselApi *CarouselManageApi) GetCarouselManageList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo huasuReq.CarouselManageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := CarouselService.GetCarouselManageInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetCarouselManagePublic 不需要鉴权的轮播图管理接口
// @Tags CarouselManage
// @Summary 不需要鉴权的轮播图管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Carousel/getCarouselManagePublic [get]
func (CarouselApi *CarouselManageApi) GetCarouselManagePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	carouselList, err := CarouselService.GetCarouselManagePublic(ctx)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(carouselList, c)
	// response.OkWithDetailed(gin.H{
	// 	carouselList,
	// }, "获取成功", c)
}
