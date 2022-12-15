import service from '@/utils/request'

// @Tags CheckIdentity
// @Summary 创建CheckIdentity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckIdentity true "创建CheckIdentity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkIdentity/createCheckIdentity [post]
export const createCheckIdentity = (data) => {
  return service({
    url: '/checkIdentity/createCheckIdentity',
    method: 'post',
    data
  })
}

// @Tags CheckIdentity
// @Summary 删除CheckIdentity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckIdentity true "删除CheckIdentity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkIdentity/deleteCheckIdentity [delete]
export const deleteCheckIdentity = (data) => {
  return service({
    url: '/checkIdentity/deleteCheckIdentity',
    method: 'delete',
    data
  })
}

// @Tags CheckIdentity
// @Summary 删除CheckIdentity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CheckIdentity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkIdentity/deleteCheckIdentity [delete]
export const deleteCheckIdentityByIds = (data) => {
  return service({
    url: '/checkIdentity/deleteCheckIdentityByIds',
    method: 'delete',
    data
  })
}

// @Tags CheckIdentity
// @Summary 更新CheckIdentity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckIdentity true "更新CheckIdentity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /checkIdentity/updateCheckIdentity [put]
export const updateCheckIdentity = (data) => {
  return service({
    url: '/checkIdentity/updateCheckIdentity',
    method: 'put',
    data
  })
}

// @Tags CheckIdentity
// @Summary 用id查询CheckIdentity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CheckIdentity true "用id查询CheckIdentity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /checkIdentity/findCheckIdentity [get]
export const findCheckIdentity = (params) => {
  return service({
    url: '/checkIdentity/findCheckIdentity',
    method: 'get',
    params
  })
}

// @Tags CheckIdentity
// @Summary 分页获取CheckIdentity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CheckIdentity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkIdentity/getCheckIdentityList [get]
export const getCheckIdentityList = (params) => {
  return service({
    url: '/checkIdentity/getCheckIdentityList',
    method: 'get',
    params
  })
}
