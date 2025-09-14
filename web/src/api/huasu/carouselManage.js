import service from '@/utils/request'
// @Tags CarouselManage
// @Summary 创建轮播图管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CarouselManage true "创建轮播图管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Carousel/createCarouselManage [post]
export const createCarouselManage = (data) => {
  return service({
    url: '/Carousel/createCarouselManage',
    method: 'post',
    data
  })
}

// @Tags CarouselManage
// @Summary 删除轮播图管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CarouselManage true "删除轮播图管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Carousel/deleteCarouselManage [delete]
export const deleteCarouselManage = (params) => {
  return service({
    url: '/Carousel/deleteCarouselManage',
    method: 'delete',
    params
  })
}

// @Tags CarouselManage
// @Summary 批量删除轮播图管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除轮播图管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Carousel/deleteCarouselManage [delete]
export const deleteCarouselManageByIds = (params) => {
  return service({
    url: '/Carousel/deleteCarouselManageByIds',
    method: 'delete',
    params
  })
}

// @Tags CarouselManage
// @Summary 更新轮播图管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CarouselManage true "更新轮播图管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Carousel/updateCarouselManage [put]
export const updateCarouselManage = (data) => {
  return service({
    url: '/Carousel/updateCarouselManage',
    method: 'put',
    data
  })
}

// @Tags CarouselManage
// @Summary 用id查询轮播图管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.CarouselManage true "用id查询轮播图管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Carousel/findCarouselManage [get]
export const findCarouselManage = (params) => {
  return service({
    url: '/Carousel/findCarouselManage',
    method: 'get',
    params
  })
}

// @Tags CarouselManage
// @Summary 分页获取轮播图管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取轮播图管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Carousel/getCarouselManageList [get]
export const getCarouselManageList = (params) => {
  return service({
    url: '/Carousel/getCarouselManageList',
    method: 'get',
    params
  })
}

// @Tags CarouselManage
// @Summary 不需要鉴权的轮播图管理接口
// @Accept application/json
// @Produce application/json
// @Param data query huasuReq.CarouselManageSearch true "分页获取轮播图管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Carousel/getCarouselManagePublic [get]
export const getCarouselManagePublic = () => {
  return service({
    url: '/Carousel/getCarouselManagePublic',
    method: 'get',
  })
}
