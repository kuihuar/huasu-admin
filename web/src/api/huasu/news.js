import service from '@/utils/request'
// @Tags News
// @Summary 创建新闻版块
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.News true "创建新闻版块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /news/createNews [post]
export const createNews = (data) => {
  return service({
    url: '/news/createNews',
    method: 'post',
    data
  })
}

// @Tags News
// @Summary 删除新闻版块
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.News true "删除新闻版块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /news/deleteNews [delete]
export const deleteNews = (params) => {
  return service({
    url: '/news/deleteNews',
    method: 'delete',
    params
  })
}

// @Tags News
// @Summary 批量删除新闻版块
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除新闻版块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /news/deleteNews [delete]
export const deleteNewsByIds = (params) => {
  return service({
    url: '/news/deleteNewsByIds',
    method: 'delete',
    params
  })
}

// @Tags News
// @Summary 更新新闻版块
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.News true "更新新闻版块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /news/updateNews [put]
export const updateNews = (data) => {
  return service({
    url: '/news/updateNews',
    method: 'put',
    data
  })
}

// @Tags News
// @Summary 用id查询新闻版块
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.News true "用id查询新闻版块"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /news/findNews [get]
export const findNews = (params) => {
  return service({
    url: '/news/findNews',
    method: 'get',
    params
  })
}

// @Tags News
// @Summary 分页获取新闻版块列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取新闻版块列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /news/getNewsList [get]
export const getNewsList = (params) => {
  return service({
    url: '/news/getNewsList',
    method: 'get',
    params
  })
}

// @Tags News
// @Summary 不需要鉴权的新闻版块接口
// @Accept application/json
// @Produce application/json
// @Param data query huasuReq.NewsSearch true "分页获取新闻版块列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /news/getNewsPublic [get]
export const getNewsPublic = () => {
  return service({
    url: '/news/getNewsPublic',
    method: 'get',
  })
}
