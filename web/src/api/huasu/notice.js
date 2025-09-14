import service from '@/utils/request'
// @Tags Notice
// @Summary 创建公告管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Notice true "创建公告管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /notice/createNotice [post]
export const createNotice = (data) => {
  return service({
    url: '/notice/createNotice',
    method: 'post',
    data
  })
}

// @Tags Notice
// @Summary 删除公告管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Notice true "删除公告管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notice/deleteNotice [delete]
export const deleteNotice = (params) => {
  return service({
    url: '/notice/deleteNotice',
    method: 'delete',
    params
  })
}

// @Tags Notice
// @Summary 批量删除公告管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公告管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notice/deleteNotice [delete]
export const deleteNoticeByIds = (params) => {
  return service({
    url: '/notice/deleteNoticeByIds',
    method: 'delete',
    params
  })
}

// @Tags Notice
// @Summary 更新公告管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Notice true "更新公告管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /notice/updateNotice [put]
export const updateNotice = (data) => {
  return service({
    url: '/notice/updateNotice',
    method: 'put',
    data
  })
}

// @Tags Notice
// @Summary 用id查询公告管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Notice true "用id查询公告管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /notice/findNotice [get]
export const findNotice = (params) => {
  return service({
    url: '/notice/findNotice',
    method: 'get',
    params
  })
}

// @Tags Notice
// @Summary 分页获取公告管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取公告管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notice/getNoticeList [get]
export const getNoticeList = (params) => {
  return service({
    url: '/notice/getNoticeList',
    method: 'get',
    params
  })
}

// @Tags Notice
// @Summary 不需要鉴权的公告管理接口
// @Accept application/json
// @Produce application/json
// @Param data query huasuReq.NoticeSearch true "分页获取公告管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /notice/getNoticePublic [get]
export const getNoticePublic = () => {
  return service({
    url: '/notice/getNoticePublic',
    method: 'get',
  })
}
