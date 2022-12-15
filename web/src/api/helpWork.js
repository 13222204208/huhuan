import service from '@/utils/request'

// @Tags HelpWork
// @Summary 创建HelpWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpWork true "创建HelpWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpWork/createHelpWork [post]
export const createHelpWork = (data) => {
  return service({
    url: '/helpWork/createHelpWork',
    method: 'post',
    data
  })
}

// @Tags HelpWork
// @Summary 删除HelpWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpWork true "删除HelpWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpWork/deleteHelpWork [delete]
export const deleteHelpWork = (data) => {
  return service({
    url: '/helpWork/deleteHelpWork',
    method: 'delete',
    data
  })
}

// @Tags HelpWork
// @Summary 删除HelpWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpWork/deleteHelpWork [delete]
export const deleteHelpWorkByIds = (data) => {
  return service({
    url: '/helpWork/deleteHelpWorkByIds',
    method: 'delete',
    data
  })
}

// @Tags HelpWork
// @Summary 更新HelpWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpWork true "更新HelpWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpWork/updateHelpWork [put]
export const updateHelpWork = (data) => {
  return service({
    url: '/helpWork/updateHelpWork',
    method: 'put',
    data
  })
}

// @Tags HelpWork
// @Summary 用id查询HelpWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HelpWork true "用id查询HelpWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpWork/findHelpWork [get]
export const findHelpWork = (params) => {
  return service({
    url: '/helpWork/findHelpWork',
    method: 'get',
    params
  })
}

// @Tags HelpWork
// @Summary 分页获取HelpWork列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HelpWork列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpWork/getHelpWorkList [get]
export const getHelpWorkList = (params) => {
  return service({
    url: '/helpWork/getHelpWorkList',
    method: 'get',
    params
  })
}
