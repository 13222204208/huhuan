import service from '@/utils/request'

// @Tags HelpIncidentally
// @Summary 创建HelpIncidentally
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpIncidentally true "创建HelpIncidentally"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpIncidentally/createHelpIncidentally [post]
export const createHelpIncidentally = (data) => {
  return service({
    url: '/helpIncidentally/createHelpIncidentally',
    method: 'post',
    data
  })
}

// @Tags HelpIncidentally
// @Summary 删除HelpIncidentally
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpIncidentally true "删除HelpIncidentally"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpIncidentally/deleteHelpIncidentally [delete]
export const deleteHelpIncidentally = (data) => {
  return service({
    url: '/helpIncidentally/deleteHelpIncidentally',
    method: 'delete',
    data
  })
}

// @Tags HelpIncidentally
// @Summary 删除HelpIncidentally
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpIncidentally"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpIncidentally/deleteHelpIncidentally [delete]
export const deleteHelpIncidentallyByIds = (data) => {
  return service({
    url: '/helpIncidentally/deleteHelpIncidentallyByIds',
    method: 'delete',
    data
  })
}

// @Tags HelpIncidentally
// @Summary 更新HelpIncidentally
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpIncidentally true "更新HelpIncidentally"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpIncidentally/updateHelpIncidentally [put]
export const updateHelpIncidentally = (data) => {
  return service({
    url: '/helpIncidentally/updateHelpIncidentally',
    method: 'put',
    data
  })
}

// @Tags HelpIncidentally
// @Summary 用id查询HelpIncidentally
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HelpIncidentally true "用id查询HelpIncidentally"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpIncidentally/findHelpIncidentally [get]
export const findHelpIncidentally = (params) => {
  return service({
    url: '/helpIncidentally/findHelpIncidentally',
    method: 'get',
    params
  })
}

// @Tags HelpIncidentally
// @Summary 分页获取HelpIncidentally列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HelpIncidentally列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpIncidentally/getHelpIncidentallyList [get]
export const getHelpIncidentallyList = (params) => {
  return service({
    url: '/helpIncidentally/getHelpIncidentallyList',
    method: 'get',
    params
  })
}
