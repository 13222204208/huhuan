import service from '@/utils/request'

// @Tags MoneyRecord
// @Summary 创建MoneyRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyRecord true "创建MoneyRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyRecord/createMoneyRecord [post]
export const createMoneyRecord = (data) => {
  return service({
    url: '/moneyRecord/createMoneyRecord',
    method: 'post',
    data
  })
}

// @Tags MoneyRecord
// @Summary 删除MoneyRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyRecord true "删除MoneyRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneyRecord/deleteMoneyRecord [delete]
export const deleteMoneyRecord = (data) => {
  return service({
    url: '/moneyRecord/deleteMoneyRecord',
    method: 'delete',
    data
  })
}

// @Tags MoneyRecord
// @Summary 删除MoneyRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoneyRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneyRecord/deleteMoneyRecord [delete]
export const deleteMoneyRecordByIds = (data) => {
  return service({
    url: '/moneyRecord/deleteMoneyRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags MoneyRecord
// @Summary 更新MoneyRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyRecord true "更新MoneyRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /moneyRecord/updateMoneyRecord [put]
export const updateMoneyRecord = (data) => {
  return service({
    url: '/moneyRecord/updateMoneyRecord',
    method: 'put',
    data
  })
}

// @Tags MoneyRecord
// @Summary 用id查询MoneyRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MoneyRecord true "用id查询MoneyRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /moneyRecord/findMoneyRecord [get]
export const findMoneyRecord = (params) => {
  return service({
    url: '/moneyRecord/findMoneyRecord',
    method: 'get',
    params
  })
}

// @Tags MoneyRecord
// @Summary 分页获取MoneyRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MoneyRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyRecord/getMoneyRecordList [get]
export const getMoneyRecordList = (params) => {
  return service({
    url: '/moneyRecord/getMoneyRecordList',
    method: 'get',
    params
  })
}
