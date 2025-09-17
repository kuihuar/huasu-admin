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

type NoticeApi struct{}

// CreateNotice 创建公告管理
// @Tags Notice
// @Summary 创建公告管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body huasu.Notice true "创建公告管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /notice/createNotice [post]
func (noticeApi *NoticeApi) CreateNotice(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var notice huasu.Notice
	err := c.ShouldBindJSON(&notice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = noticeService.CreateNotice(ctx, &notice)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteNotice 删除公告管理
// @Tags Notice
// @Summary 删除公告管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body huasu.Notice true "删除公告管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /notice/deleteNotice [delete]
func (noticeApi *NoticeApi) DeleteNotice(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := noticeService.DeleteNotice(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteNoticeByIds 批量删除公告管理
// @Tags Notice
// @Summary 批量删除公告管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /notice/deleteNoticeByIds [delete]
func (noticeApi *NoticeApi) DeleteNoticeByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := noticeService.DeleteNoticeByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateNotice 更新公告管理
// @Tags Notice
// @Summary 更新公告管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body huasu.Notice true "更新公告管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /notice/updateNotice [put]
func (noticeApi *NoticeApi) UpdateNotice(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var notice huasu.Notice
	err := c.ShouldBindJSON(&notice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = noticeService.UpdateNotice(ctx, notice)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindNotice 用id查询公告管理
// @Tags Notice
// @Summary 用id查询公告管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询公告管理"
// @Success 200 {object} response.Response{data=huasu.Notice,msg=string} "查询成功"
// @Router /notice/findNotice [get]
func (noticeApi *NoticeApi) FindNotice(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	renotice, err := noticeService.GetNotice(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(renotice, c)
}

// GetNoticeList 分页获取公告管理列表
// @Tags Notice
// @Summary 分页获取公告管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query huasuReq.NoticeSearch true "分页获取公告管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /notice/getNoticeList [get]
func (noticeApi *NoticeApi) GetNoticeList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo huasuReq.NoticeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := noticeService.GetNoticeInfoList(ctx, pageInfo)
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

// GetNoticePublic 不需要鉴权的公告管理接口
// @Tags Notice
// @Summary 不需要鉴权的公告管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /notice/getNoticePublic [get]
func (noticeApi *NoticeApi) GetNoticePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	notices, err := noticeService.GetNoticePublic(ctx)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(notices, c)
}

// GetNewsPublic 不需要鉴权的公告接口
// @Tags News
// @Summary 不需要鉴权的公告接口
// @Accept application/json
// @Produce application/json
// @Param data query huasuReq.NoticeSearch true "分页获取公告列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /notice/GetNoticePublicList [get]
func (noticeApi *NoticeApi) GetNoticePublicList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	search := huasuReq.NoticeSearch{
		PageInfo: request.PageInfo{
			Page:     page,
			PageSize: pageSize,
		},
	}
	notices, total, err := noticeService.GetNoticePublicList(ctx, search)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{
		"info":  notices,
		"total": total,
	}, "获取成功", c)
}

// GetNoticePublicDetail 不需要鉴权的公告接口
// @Tags News
// @Summary 不需要鉴权的公告接口
// @Accept application/json
// @Produce application/json
// @Param id query string true "公告ID"
// @Success 200 {object} response.Response{data=huasu.Notice,msg=string} "获取成功"
// @Router /notice/GetNoticePublicDetail [get]
func (noticeApi *NoticeApi) GetNoticePublicDetail(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	id := c.DefaultQuery("id", "1")

	notice, err := noticeService.GetNoticePublicDetail(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(notice, c)
}
