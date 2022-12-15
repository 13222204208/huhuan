import service from '@/utils/request'

// @Tags HelpRepair
// @Summary 创建HelpRepair
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpRepair true "创建HelpRepair"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpRepair/createHelpRepair [post]
export const createHelpRepair = (data) => {
  return service({
    url: '/helpRepair/createHelpRepair',
    method: 'post',
    data
  })
}

// @Tags HelpRepair
// @Summary 删除HelpRepair
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpRepair true "删除HelpRepair"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpRepair/deleteHelpRepair [delete]
export const deleteHelpRepair = (data) => {
  return service({
    url: '/helpRepair/deleteHelpRepair',
    method: 'delete',
    data
  })
}

// @Tags HelpRepair
// @Summary 删除HelpRepair
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpRepair"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpRepair/deleteHelpRepair [delete]
export const deleteHelpRepairByIds = (data) => {
  return service({
    url: '/helpRepair/deleteHelpRepairByIds',
    method: 'delete',
    data
  })
}

// @Tags HelpRepair
// @Summary 更新HelpRepair
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpRepair true "更新HelpRepair"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpRepair/updateHelpRepair [put]
export const updateHelpRepair = (data) => {
  return service({
    url: '/helpRepair/updateHelpRepair',
    method: 'put',
    data
  })
}

// @Tags HelpRepair
// @Summary 用id查询HelpRepair
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HelpRepair true "用id查询HelpRepair"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpRepair/findHelpRepair [get]
export const findHelpRepair = (params) => {
  return service({
    url: '/helpRepair/findHelpRepair',
    method: 'get',
    params
  })
}

// @Tags HelpRepair
// @Summary 分页获取HelpRepair列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HelpRepair列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpRepair/getHelpRepairList [get]
export const getHelpRepairList = (params) => {
  return service({
    url: '/helpRepair/getHelpRepairList',
    method: 'get',
    params
  })
}
