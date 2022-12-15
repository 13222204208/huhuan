import service from '@/utils/request'

// @Tags HelpCarpool
// @Summary 创建HelpCarpool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpCarpool true "创建HelpCarpool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpCarpool/createHelpCarpool [post]
export const createHelpCarpool = (data) => {
  return service({
    url: '/helpCarpool/createHelpCarpool',
    method: 'post',
    data
  })
}

// @Tags HelpCarpool
// @Summary 删除HelpCarpool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpCarpool true "删除HelpCarpool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpCarpool/deleteHelpCarpool [delete]
export const deleteHelpCarpool = (data) => {
  return service({
    url: '/helpCarpool/deleteHelpCarpool',
    method: 'delete',
    data
  })
}

// @Tags HelpCarpool
// @Summary 删除HelpCarpool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpCarpool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpCarpool/deleteHelpCarpool [delete]
export const deleteHelpCarpoolByIds = (data) => {
  return service({
    url: '/helpCarpool/deleteHelpCarpoolByIds',
    method: 'delete',
    data
  })
}

// @Tags HelpCarpool
// @Summary 更新HelpCarpool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpCarpool true "更新HelpCarpool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpCarpool/updateHelpCarpool [put]
export const updateHelpCarpool = (data) => {
  return service({
    url: '/helpCarpool/updateHelpCarpool',
    method: 'put',
    data
  })
}

// @Tags HelpCarpool
// @Summary 用id查询HelpCarpool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HelpCarpool true "用id查询HelpCarpool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpCarpool/findHelpCarpool [get]
export const findHelpCarpool = (params) => {
  return service({
    url: '/helpCarpool/findHelpCarpool',
    method: 'get',
    params
  })
}

// @Tags HelpCarpool
// @Summary 分页获取HelpCarpool列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HelpCarpool列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpCarpool/getHelpCarpoolList [get]
export const getHelpCarpoolList = (params) => {
  return service({
    url: '/helpCarpool/getHelpCarpoolList',
    method: 'get',
    params
  })
}
