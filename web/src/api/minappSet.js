import service from '@/utils/request'

// @Tags MinappSet
// @Summary 创建MinappSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MinappSet true "创建MinappSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /minappSet/createMinappSet [post]
export const createMinappSet = (data) => {
  return service({
    url: '/minappSet/createMinappSet',
    method: 'post',
    data
  })
}

// @Tags MinappSet
// @Summary 删除MinappSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MinappSet true "删除MinappSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /minappSet/deleteMinappSet [delete]
export const deleteMinappSet = (data) => {
  return service({
    url: '/minappSet/deleteMinappSet',
    method: 'delete',
    data
  })
}

// @Tags MinappSet
// @Summary 删除MinappSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MinappSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /minappSet/deleteMinappSet [delete]
export const deleteMinappSetByIds = (data) => {
  return service({
    url: '/minappSet/deleteMinappSetByIds',
    method: 'delete',
    data
  })
}

// @Tags MinappSet
// @Summary 更新MinappSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MinappSet true "更新MinappSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /minappSet/updateMinappSet [put]
export const updateMinappSet = (data) => {
  return service({
    url: '/minappSet/updateMinappSet',
    method: 'put',
    data
  })
}

// @Tags MinappSet
// @Summary 用id查询MinappSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MinappSet true "用id查询MinappSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /minappSet/findMinappSet [get]
export const findMinappSet = (params) => {
  return service({
    url: '/minappSet/findMinappSet',
    method: 'get',
    params
  })
}

// @Tags MinappSet
// @Summary 分页获取MinappSet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MinappSet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /minappSet/getMinappSetList [get]
export const getMinappSetList = (params) => {
  return service({
    url: '/minappSet/getMinappSetList',
    method: 'get',
    params
  })
}
