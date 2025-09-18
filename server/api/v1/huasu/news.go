package huasu

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuihuar/huasu-admin/server/global"
	"github.com/kuihuar/huasu-admin/server/model/common/request"
	"github.com/kuihuar/huasu-admin/server/model/common/response"
	"github.com/kuihuar/huasu-admin/server/model/huasu"
	huasuReq "github.com/kuihuar/huasu-admin/server/model/huasu/request"
	"go.uber.org/zap"
)

type NewsApi struct{}

// CreateNews 创建新闻版块
// @Tags News
// @Summary 创建新闻版块
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body huasu.News true "创建新闻版块"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /news/createNews [post]
func (newsApi *NewsApi) CreateNews(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var news huasu.News
	err := c.ShouldBindJSON(&news)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = newsService.CreateNews(ctx, &news)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteNews 删除新闻版块
// @Tags News
// @Summary 删除新闻版块
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body huasu.News true "删除新闻版块"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /news/deleteNews [delete]
func (newsApi *NewsApi) DeleteNews(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := newsService.DeleteNews(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteNewsByIds 批量删除新闻版块
// @Tags News
// @Summary 批量删除新闻版块
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /news/deleteNewsByIds [delete]
func (newsApi *NewsApi) DeleteNewsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := newsService.DeleteNewsByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateNews 更新新闻版块
// @Tags News
// @Summary 更新新闻版块
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body huasu.News true "更新新闻版块"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /news/updateNews [put]
func (newsApi *NewsApi) UpdateNews(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var news huasu.News
	err := c.ShouldBindJSON(&news)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = newsService.UpdateNews(ctx, news)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindNews 用id查询新闻版块
// @Tags News
// @Summary 用id查询新闻版块
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询新闻版块"
// @Success 200 {object} response.Response{data=huasu.News,msg=string} "查询成功"
// @Router /news/findNews [get]
func (newsApi *NewsApi) FindNews(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	renews, err := newsService.GetNews(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(renews, c)
}

// GetNewsList 分页获取新闻版块列表
// @Tags News
// @Summary 分页获取新闻版块列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query huasuReq.NewsSearch true "分页获取新闻版块列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /news/getNewsList [get]
func (newsApi *NewsApi) GetNewsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo huasuReq.NewsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := newsService.GetNewsInfoList(ctx, pageInfo)
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

// GetNewsPublic 不需要鉴权的新闻版块接口
// @Tags News
// @Summary 不需要鉴权的新闻版块接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /news/getNewsPublic [get]
func (newsApi *NewsApi) GetNewsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	cate := c.DefaultQuery("category", "")

	search := huasuReq.NewsSearch{
		PageInfo: request.PageInfo{
			Page:     page,
			PageSize: pageSize,
		},
	}
	if cate != "" {
		search.Category = cate
	}

	newsList, total, err := newsService.GetNewsPublic(ctx, search)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{
		"info":  newsList,
		"total": total,
	}, "获取成功", c)
}
