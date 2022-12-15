import service from '@/utils/request'

// @Tags Withdraw
// @Summary 创建Withdraw
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Withdraw true "创建Withdraw"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /withdraw/createWithdraw [post]
export const createWithdraw = (data) => {
  return service({
    url: '/withdraw/createWithdraw',
    method: 'post',
    data
  })
}

// @Tags Withdraw
// @Summary 删除Withdraw
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Withdraw true "删除Withdraw"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /withdraw/deleteWithdraw [delete]
export const deleteWithdraw = (data) => {
  return service({
    url: '/withdraw/deleteWithdraw',
    method: 'delete',
    data
  })
}

// @Tags Withdraw
// @Summary 删除Withdraw
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Withdraw"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /withdraw/deleteWithdraw [delete]
export const deleteWithdrawByIds = (data) => {
  return service({
    url: '/withdraw/deleteWithdrawByIds',
    method: 'delete',
    data
  })
}

// @Tags Withdraw
// @Summary 更新Withdraw
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Withdraw true "更新Withdraw"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /withdraw/updateWithdraw [put]
export const updateWithdraw = (data) => {
  return service({
    url: '/withdraw/updateWithdraw',
    method: 'put',
    data
  })
}

// @Tags Withdraw
// @Summary 用id查询Withdraw
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Withdraw true "用id查询Withdraw"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /withdraw/findWithdraw [get]
export const findWithdraw = (params) => {
  return service({
    url: '/withdraw/findWithdraw',
    method: 'get',
    params
  })
}

// @Tags Withdraw
// @Summary 分页获取Withdraw列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Withdraw列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /withdraw/getWithdrawList [get]
export const getWithdrawList = (params) => {
  return service({
    url: '/withdraw/getWithdrawList',
    method: 'get',
    params
  })
}
