import service from '@/utils/request'
// @Tags ConstructionManage
// @Summary 创建工程管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ConstructionManage true "创建工程管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /constructionManage/createConstructionManage [post]
export const createConstructionManage = (data) => {
  return service({
    url: '/constructionManage/createConstructionManage',
    method: 'post',
    data
  })
}

// @Tags ConstructionManage
// @Summary 删除工程管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ConstructionManage true "删除工程管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /constructionManage/deleteConstructionManage [delete]
export const deleteConstructionManage = (params) => {
  return service({
    url: '/constructionManage/deleteConstructionManage',
    method: 'delete',
    params
  })
}

// @Tags ConstructionManage
// @Summary 批量删除工程管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工程管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /constructionManage/deleteConstructionManage [delete]
export const deleteConstructionManageByIds = (params) => {
  return service({
    url: '/constructionManage/deleteConstructionManageByIds',
    method: 'delete',
    params
  })
}

// @Tags ConstructionManage
// @Summary 更新工程管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ConstructionManage true "更新工程管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /constructionManage/updateConstructionManage [put]
export const updateConstructionManage = (data) => {
  return service({
    url: '/constructionManage/updateConstructionManage',
    method: 'put',
    data
  })
}

// @Tags ConstructionManage
// @Summary 用id查询工程管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ConstructionManage true "用id查询工程管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /constructionManage/findConstructionManage [get]
export const findConstructionManage = (params) => {
  return service({
    url: '/constructionManage/findConstructionManage',
    method: 'get',
    params
  })
}

// @Tags ConstructionManage
// @Summary 分页获取工程管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工程管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /constructionManage/getConstructionManageList [get]
export const getConstructionManageList = (params) => {
  return service({
    url: '/constructionManage/getConstructionManageList',
    method: 'get',
    params
  })
}

// @Tags ConstructionManage
// @Summary 不需要鉴权的工程管理接口
// @Accept application/json
// @Produce application/json
// @Param data query huasuReq.ConstructionManageSearch true "分页获取工程管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /constructionManage/getConstructionManagePublic [get]
export const getConstructionManagePublic = () => {
  return service({
    url: '/constructionManage/getConstructionManagePublic',
    method: 'get',
  })
}
