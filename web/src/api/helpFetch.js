import service from '@/utils/request'

// @Tags HelpFetch
// @Summary 创建HelpFetch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpFetch true "创建HelpFetch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpFetch/createHelpFetch [post]
export const createHelpFetch = (data) => {
  return service({
    url: '/helpFetch/createHelpFetch',
    method: 'post',
    data
  })
}

// @Tags HelpFetch
// @Summary 删除HelpFetch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpFetch true "删除HelpFetch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpFetch/deleteHelpFetch [delete]
export const deleteHelpFetch = (data) => {
  return service({
    url: '/helpFetch/deleteHelpFetch',
    method: 'delete',
    data
  })
}

// @Tags HelpFetch
// @Summary 删除HelpFetch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HelpFetch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /helpFetch/deleteHelpFetch [delete]
export const deleteHelpFetchByIds = (data) => {
  return service({
    url: '/helpFetch/deleteHelpFetchByIds',
    method: 'delete',
    data
  })
}

// @Tags HelpFetch
// @Summary 更新HelpFetch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelpFetch true "更新HelpFetch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /helpFetch/updateHelpFetch [put]
export const updateHelpFetch = (data) => {
  return service({
    url: '/helpFetch/updateHelpFetch',
    method: 'put',
    data
  })
}

// @Tags HelpFetch
// @Summary 用id查询HelpFetch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HelpFetch true "用id查询HelpFetch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /helpFetch/findHelpFetch [get]
export const findHelpFetch = (params) => {
  return service({
    url: '/helpFetch/findHelpFetch',
    method: 'get',
    params
  })
}

// @Tags HelpFetch
// @Summary 分页获取HelpFetch列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HelpFetch列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /helpFetch/getHelpFetchList [get]
export const getHelpFetchList = (params) => {
  return service({
    url: '/helpFetch/getHelpFetchList',
    method: 'get',
    params
  })
}
