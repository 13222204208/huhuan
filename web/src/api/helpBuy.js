import service from '@/utils/request'

// @Tags HelpBuy
// @Summary 创建HelpBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpBuy true "创建HelpBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpBuy/createHelpBuy [post]
export const createHelpBuy = (data) => {
  return service({
    url: '/helpBuy/createHelpBuy',
    method: 'post',
    data
  })
}

// @Tags HelpBuy
// @Summary 删除HelpBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpBuy true "删除HelpBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpBuy/deleteHelpBuy [delete]
export const deleteHelpBuy = (data) => {
  return service({
    url: '/helpBuy/deleteHelpBuy',
    method: 'delete',
    data
  })
}

// @Tags HelpBuy
// @Summary 删除HelpBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpBuy/deleteHelpBuy [delete]
export const deleteHelpBuyByIds = (data) => {
  return service({
    url: '/helpBuy/deleteHelpBuyByIds',
    method: 'delete',
    data
  })
}

// @Tags HelpBuy
// @Summary 更新HelpBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpBuy true "更新HelpBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpBuy/updateHelpBuy [put]
export const updateHelpBuy = (data) => {
  return service({
    url: '/helpBuy/updateHelpBuy',
    method: 'put',
    data
  })
}

// @Tags HelpBuy
// @Summary 用id查询HelpBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HelpBuy true "用id查询HelpBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpBuy/findHelpBuy [get]
export const findHelpBuy = (params) => {
  return service({
    url: '/helpBuy/findHelpBuy',
    method: 'get',
    params
  })
}

// @Tags HelpBuy
// @Summary 分页获取HelpBuy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HelpBuy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpBuy/getHelpBuyList [get]
export const getHelpBuyList = (params) => {
  return service({
    url: '/helpBuy/getHelpBuyList',
    method: 'get',
    params
  })
}
