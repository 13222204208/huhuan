import service from '@/utils/request'

// @Tags MinappUser
// @Summary 创建MinappUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MinappUser true "创建MinappUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /minappUser/createMinappUser [post]
export const createMinappUser = (data) => {
  return service({
    url: '/minappUser/createMinappUser',
    method: 'post',
    data
  })
}

// @Tags MinappUser
// @Summary 删除MinappUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MinappUser true "删除MinappUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /minappUser/deleteMinappUser [delete]
export const deleteMinappUser = (data) => {
  return service({
    url: '/minappUser/deleteMinappUser',
    method: 'delete',
    data
  })
}

// @Tags MinappUser
// @Summary 删除MinappUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MinappUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /minappUser/deleteMinappUser [delete]
export const deleteMinappUserByIds = (data) => {
  return service({
    url: '/minappUser/deleteMinappUserByIds',
    method: 'delete',
    data
  })
}

// @Tags MinappUser
// @Summary 更新MinappUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MinappUser true "更新MinappUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /minappUser/updateMinappUser [put]
export const updateMinappUser = (data) => {
  return service({
    url: '/minappUser/updateMinappUser',
    method: 'put',
    data
  })
}

// @Tags MinappUser
// @Summary 用id查询MinappUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MinappUser true "用id查询MinappUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /minappUser/findMinappUser [get]
export const findMinappUser = (params) => {
  return service({
    url: '/minappUser/findMinappUser',
    method: 'get',
    params
  })
}

// @Tags MinappUser
// @Summary 分页获取MinappUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MinappUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /minappUser/getMinappUserList [get]
export const getMinappUserList = (params) => {
  return service({
    url: '/minappUser/getMinappUserList',
    method: 'get',
    params
  })
}

//获取新增用户每天的数量
export const getNewUserCountList = (params) => {
  return service({
    url: '/minappUser/getNewUserCountList',
    method: 'get',
    params
  })
}