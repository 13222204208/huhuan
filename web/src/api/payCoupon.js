import service from '@/utils/request'

// @Tags PayCoupon
// @Summary 创建PayCoupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayCoupon true "创建PayCoupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payCoupon/createPayCoupon [post]
export const createPayCoupon = (data) => {
  return service({
    url: '/payCoupon/createPayCoupon',
    method: 'post',
    data
  })
}

// @Tags PayCoupon
// @Summary 删除PayCoupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayCoupon true "删除PayCoupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payCoupon/deletePayCoupon [delete]
export const deletePayCoupon = (data) => {
  return service({
    url: '/payCoupon/deletePayCoupon',
    method: 'delete',
    data
  })
}

// @Tags PayCoupon
// @Summary 删除PayCoupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayCoupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payCoupon/deletePayCoupon [delete]
export const deletePayCouponByIds = (data) => {
  return service({
    url: '/payCoupon/deletePayCouponByIds',
    method: 'delete',
    data
  })
}

// @Tags PayCoupon
// @Summary 更新PayCoupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayCoupon true "更新PayCoupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payCoupon/updatePayCoupon [put]
export const updatePayCoupon = (data) => {
  return service({
    url: '/payCoupon/updatePayCoupon',
    method: 'put',
    data
  })
}

// @Tags PayCoupon
// @Summary 用id查询PayCoupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PayCoupon true "用id查询PayCoupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payCoupon/findPayCoupon [get]
export const findPayCoupon = (params) => {
  return service({
    url: '/payCoupon/findPayCoupon',
    method: 'get',
    params
  })
}

// @Tags PayCoupon
// @Summary 分页获取PayCoupon列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PayCoupon列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payCoupon/getPayCouponList [get]
export const getPayCouponList = (params) => {
  return service({
    url: '/payCoupon/getPayCouponList',
    method: 'get',
    params
  })
}
