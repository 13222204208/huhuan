import service from '@/utils/request'

// @Tags CreditRating
// @Summary 创建CreditRating
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CreditRating true "创建CreditRating"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /creditRating/createCreditRating [post]
export const createCreditRating = (data) => {
  return service({
    url: '/creditRating/createCreditRating',
    method: 'post',
    data
  })
}

// @Tags CreditRating
// @Summary 删除CreditRating
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CreditRating true "删除CreditRating"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /creditRating/deleteCreditRating [delete]
export const deleteCreditRating = (data) => {
  return service({
    url: '/creditRating/deleteCreditRating',
    method: 'delete',
    data
  })
}

// @Tags CreditRating
// @Summary 删除CreditRating
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CreditRating"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /creditRating/deleteCreditRating [delete]
export const deleteCreditRatingByIds = (data) => {
  return service({
    url: '/creditRating/deleteCreditRatingByIds',
    method: 'delete',
    data
  })
}

// @Tags CreditRating
// @Summary 更新CreditRating
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CreditRating true "更新CreditRating"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /creditRating/updateCreditRating [put]
export const updateCreditRating = (data) => {
  return service({
    url: '/creditRating/updateCreditRating',
    method: 'put',
    data
  })
}

// @Tags CreditRating
// @Summary 用id查询CreditRating
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CreditRating true "用id查询CreditRating"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /creditRating/findCreditRating [get]
export const findCreditRating = (params) => {
  return service({
    url: '/creditRating/findCreditRating',
    method: 'get',
    params
  })
}

// @Tags CreditRating
// @Summary 分页获取CreditRating列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CreditRating列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /creditRating/getCreditRatingList [get]
export const getCreditRatingList = (params) => {
  return service({
    url: '/creditRating/getCreditRatingList',
    method: 'get',
    params
  })
}
