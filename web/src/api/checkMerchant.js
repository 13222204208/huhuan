import service from '@/utils/request'

// @Tags CheckMerchant
// @Summary 创建CheckMerchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckMerchant true "创建CheckMerchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkMerchant/createCheckMerchant [post]
export const createCheckMerchant = (data) => {
  return service({
    url: '/checkMerchant/createCheckMerchant',
    method: 'post',
    data
  })
}

// @Tags CheckMerchant
// @Summary 删除CheckMerchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckMerchant true "删除CheckMerchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkMerchant/deleteCheckMerchant [delete]
export const deleteCheckMerchant = (data) => {
  return service({
    url: '/checkMerchant/deleteCheckMerchant',
    method: 'delete',
    data
  })
}

// @Tags CheckMerchant
// @Summary 删除CheckMerchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CheckMerchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkMerchant/deleteCheckMerchant [delete]
export const deleteCheckMerchantByIds = (data) => {
  return service({
    url: '/checkMerchant/deleteCheckMerchantByIds',
    method: 'delete',
    data
  })
}

// @Tags CheckMerchant
// @Summary 更新CheckMerchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckMerchant true "更新CheckMerchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /checkMerchant/updateCheckMerchant [put]
export const updateCheckMerchant = (data) => {
  return service({
    url: '/checkMerchant/updateCheckMerchant',
    method: 'put',
    data
  })
}

// @Tags CheckMerchant
// @Summary 用id查询CheckMerchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CheckMerchant true "用id查询CheckMerchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /checkMerchant/findCheckMerchant [get]
export const findCheckMerchant = (params) => {
  return service({
    url: '/checkMerchant/findCheckMerchant',
    method: 'get',
    params
  })
}

// @Tags CheckMerchant
// @Summary 分页获取CheckMerchant列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CheckMerchant列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkMerchant/getCheckMerchantList [get]
export const getCheckMerchantList = (params) => {
  return service({
    url: '/checkMerchant/getCheckMerchantList',
    method: 'get',
    params
  })
}
