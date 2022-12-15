import service from '@/utils/request'

// @Tags Commission
// @Summary 创建Commission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Commission true "创建Commission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /commission/createCommission [post]
export const createCommission = (data) => {
  return service({
    url: '/commission/createCommission',
    method: 'post',
    data
  })
}

// @Tags Commission
// @Summary 删除Commission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Commission true "删除Commission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commission/deleteCommission [delete]
export const deleteCommission = (data) => {
  return service({
    url: '/commission/deleteCommission',
    method: 'delete',
    data
  })
}

// @Tags Commission
// @Summary 删除Commission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Commission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commission/deleteCommission [delete]
export const deleteCommissionByIds = (data) => {
  return service({
    url: '/commission/deleteCommissionByIds',
    method: 'delete',
    data
  })
}

// @Tags Commission
// @Summary 更新Commission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Commission true "更新Commission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /commission/updateCommission [put]
export const updateCommission = (data) => {
  return service({
    url: '/commission/updateCommission',
    method: 'put',
    data
  })
}

// @Tags Commission
// @Summary 用id查询Commission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Commission true "用id查询Commission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /commission/findCommission [get]
export const findCommission = (params) => {
  return service({
    url: '/commission/findCommission',
    method: 'get',
    params
  })
}

// @Tags Commission
// @Summary 分页获取Commission列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Commission列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /commission/getCommissionList [get]
export const getCommissionList = (params) => {
  return service({
    url: '/commission/getCommissionList',
    method: 'get',
    params
  })
}
