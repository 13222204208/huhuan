import service from '@/utils/request'

// @Tags Coupon
// @Summary 创建Coupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coupon true "创建Coupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coupon/createCoupon [post]
export const createCoupon = (data) => {
  return service({
    url: '/coupon/createCoupon',
    method: 'post',
    data
  })
}

// @Tags Coupon
// @Summary 删除Coupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coupon true "删除Coupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /coupon/deleteCoupon [delete]
export const deleteCoupon = (data) => {
  return service({
    url: '/coupon/deleteCoupon',
    method: 'delete',
    data
  })
}

// @Tags Coupon
// @Summary 删除Coupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Coupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /coupon/deleteCoupon [delete]
export const deleteCouponByIds = (data) => {
  return service({
    url: '/coupon/deleteCouponByIds',
    method: 'delete',
    data
  })
}

// @Tags Coupon
// @Summary 更新Coupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coupon true "更新Coupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /coupon/updateCoupon [put]
export const updateCoupon = (data) => {
  return service({
    url: '/coupon/updateCoupon',
    method: 'put',
    data
  })
}

// @Tags Coupon
// @Summary 用id查询Coupon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Coupon true "用id查询Coupon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /coupon/findCoupon [get]
export const findCoupon = (params) => {
  return service({
    url: '/coupon/findCoupon',
    method: 'get',
    params
  })
}

// @Tags Coupon
// @Summary 分页获取Coupon列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Coupon列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coupon/getCouponList [get]
export const getCouponList = (params) => {
  return service({
    url: '/coupon/getCouponList',
    method: 'get',
    params
  })
}

export const getAllCouponTitle = () => {
  return service({
    url:'/coupon/getAllCoupons',
    method: 'get'
  })
}
