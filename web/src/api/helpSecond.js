import service from '@/utils/request'

// @Tags HelpSecond
// @Summary 创建HelpSecond
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpSecond true "创建HelpSecond"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpSecond/createHelpSecond [post]
export const createHelpSecond = (data) => {
  return service({
    url: '/helpSecond/createHelpSecond',
    method: 'post',
    data
  })
}

// @Tags HelpSecond
// @Summary 删除HelpSecond
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpSecond true "删除HelpSecond"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpSecond/deleteHelpSecond [delete]
export const deleteHelpSecond = (data) => {
  return service({
    url: '/helpSecond/deleteHelpSecond',
    method: 'delete',
    data
  })
}

// @Tags HelpSecond
// @Summary 删除HelpSecond
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpSecond"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpSecond/deleteHelpSecond [delete]
export const deleteHelpSecondByIds = (data) => {
  return service({
    url: '/helpSecond/deleteHelpSecondByIds',
    method: 'delete',
    data
  })
}

// @Tags HelpSecond
// @Summary 更新HelpSecond
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpSecond true "更新HelpSecond"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpSecond/updateHelpSecond [put]
export const updateHelpSecond = (data) => {
  return service({
    url: '/helpSecond/updateHelpSecond',
    method: 'put',
    data
  })
}

// @Tags HelpSecond
// @Summary 用id查询HelpSecond
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HelpSecond true "用id查询HelpSecond"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpSecond/findHelpSecond [get]
export const findHelpSecond = (params) => {
  return service({
    url: '/helpSecond/findHelpSecond',
    method: 'get',
    params
  })
}

// @Tags HelpSecond
// @Summary 分页获取HelpSecond列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HelpSecond列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpSecond/getHelpSecondList [get]
export const getHelpSecondList = (params) => {
  return service({
    url: '/helpSecond/getHelpSecondList',
    method: 'get',
    params
  })
}
