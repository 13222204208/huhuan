import service from '@/utils/request'

// @Tags HelpGroupon
// @Summary 创建HelpGroupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpGroupon true "创建HelpGroupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpGroupon/createHelpGroupon [post]
export const createHelpGroupon = (data) => {
  return service({
    url: '/helpGroupon/createHelpGroupon',
    method: 'post',
    data
  })
}

// @Tags HelpGroupon
// @Summary 删除HelpGroupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpGroupon true "删除HelpGroupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpGroupon/deleteHelpGroupon [delete]
export const deleteHelpGroupon = (data) => {
  return service({
    url: '/helpGroupon/deleteHelpGroupon',
    method: 'delete',
    data
  })
}

// @Tags HelpGroupon
// @Summary 删除HelpGroupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpGroupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpGroupon/deleteHelpGroupon [delete]
export const deleteHelpGrouponByIds = (data) => {
  return service({
    url: '/helpGroupon/deleteHelpGrouponByIds',
    method: 'delete',
    data
  })
}

// @Tags HelpGroupon
// @Summary 更新HelpGroupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpGroupon true "更新HelpGroupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpGroupon/updateHelpGroupon [put]
export const updateHelpGroupon = (data) => {
  return service({
    url: '/helpGroupon/updateHelpGroupon',
    method: 'put',
    data
  })
}

// @Tags HelpGroupon
// @Summary 用id查询HelpGroupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HelpGroupon true "用id查询HelpGroupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpGroupon/findHelpGroupon [get]
export const findHelpGroupon = (params) => {
  return service({
    url: '/helpGroupon/findHelpGroupon',
    method: 'get',
    params
  })
}

// @Tags HelpGroupon
// @Summary 分页获取HelpGroupon列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HelpGroupon列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpGroupon/getHelpGrouponList [get]
export const getHelpGrouponList = (params) => {
  return service({
    url: '/helpGroupon/getHelpGrouponList',
    method: 'get',
    params
  })
}
