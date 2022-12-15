import service from '@/utils/request'

// @Tags SystemMessage
// @Summary 创建SystemMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemMessage true "创建SystemMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemMessage/createSystemMessage [post]
export const createSystemMessage = (data) => {
  return service({
    url: '/systemMessage/createSystemMessage',
    method: 'post',
    data
  })
}

// @Tags SystemMessage
// @Summary 删除SystemMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemMessage true "删除SystemMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /systemMessage/deleteSystemMessage [delete]
export const deleteSystemMessage = (data) => {
  return service({
    url: '/systemMessage/deleteSystemMessage',
    method: 'delete',
    data
  })
}

// @Tags SystemMessage
// @Summary 删除SystemMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SystemMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /systemMessage/deleteSystemMessage [delete]
export const deleteSystemMessageByIds = (data) => {
  return service({
    url: '/systemMessage/deleteSystemMessageByIds',
    method: 'delete',
    data
  })
}

// @Tags SystemMessage
// @Summary 更新SystemMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemMessage true "更新SystemMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /systemMessage/updateSystemMessage [put]
export const updateSystemMessage = (data) => {
  return service({
    url: '/systemMessage/updateSystemMessage',
    method: 'put',
    data
  })
}

// @Tags SystemMessage
// @Summary 用id查询SystemMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SystemMessage true "用id查询SystemMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /systemMessage/findSystemMessage [get]
export const findSystemMessage = (params) => {
  return service({
    url: '/systemMessage/findSystemMessage',
    method: 'get',
    params
  })
}

// @Tags SystemMessage
// @Summary 分页获取SystemMessage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SystemMessage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemMessage/getSystemMessageList [get]
export const getSystemMessageList = (params) => {
  return service({
    url: '/systemMessage/getSystemMessageList',
    method: 'get',
    params
  })
}
