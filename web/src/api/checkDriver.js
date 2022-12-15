import service from '@/utils/request'

// @Tags CheckDriver
// @Summary 创建CheckDriver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckDriver true "创建CheckDriver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkDriver/createCheckDriver [post]
export const createCheckDriver = (data) => {
  return service({
    url: '/checkDriver/createCheckDriver',
    method: 'post',
    data
  })
}

// @Tags CheckDriver
// @Summary 删除CheckDriver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckDriver true "删除CheckDriver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkDriver/deleteCheckDriver [delete]
export const deleteCheckDriver = (data) => {
  return service({
    url: '/checkDriver/deleteCheckDriver',
    method: 'delete',
    data
  })
}

// @Tags CheckDriver
// @Summary 删除CheckDriver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CheckDriver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkDriver/deleteCheckDriver [delete]
export const deleteCheckDriverByIds = (data) => {
  return service({
    url: '/checkDriver/deleteCheckDriverByIds',
    method: 'delete',
    data
  })
}

// @Tags CheckDriver
// @Summary 更新CheckDriver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckDriver true "更新CheckDriver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /checkDriver/updateCheckDriver [put]
export const updateCheckDriver = (data) => {
  return service({
    url: '/checkDriver/updateCheckDriver',
    method: 'put',
    data
  })
}

// @Tags CheckDriver
// @Summary 用id查询CheckDriver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CheckDriver true "用id查询CheckDriver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /checkDriver/findCheckDriver [get]
export const findCheckDriver = (params) => {
  return service({
    url: '/checkDriver/findCheckDriver',
    method: 'get',
    params
  })
}

// @Tags CheckDriver
// @Summary 分页获取CheckDriver列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CheckDriver列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkDriver/getCheckDriverList [get]
export const getCheckDriverList = (params) => {
  return service({
    url: '/checkDriver/getCheckDriverList',
    method: 'get',
    params
  })
}
