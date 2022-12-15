import service from '@/utils/request'

// @Tags Agreement
// @Summary 创建Agreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Agreement true "创建Agreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agreement/createAgreement [post]
export const createAgreement = (data) => {
  return service({
    url: '/agreement/createAgreement',
    method: 'post',
    data
  })
}

// @Tags Agreement
// @Summary 删除Agreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Agreement true "删除Agreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /agreement/deleteAgreement [delete]
export const deleteAgreement = (data) => {
  return service({
    url: '/agreement/deleteAgreement',
    method: 'delete',
    data
  })
}

// @Tags Agreement
// @Summary 删除Agreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Agreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /agreement/deleteAgreement [delete]
export const deleteAgreementByIds = (data) => {
  return service({
    url: '/agreement/deleteAgreementByIds',
    method: 'delete',
    data
  })
}

// @Tags Agreement
// @Summary 更新Agreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Agreement true "更新Agreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /agreement/updateAgreement [put]
export const updateAgreement = (data) => {
  return service({
    url: '/agreement/updateAgreement',
    method: 'put',
    data
  })
}

// @Tags Agreement
// @Summary 用id查询Agreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Agreement true "用id查询Agreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /agreement/findAgreement [get]
export const findAgreement = (params) => {
  return service({
    url: '/agreement/findAgreement',
    method: 'get',
    params
  })
}

// @Tags Agreement
// @Summary 分页获取Agreement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Agreement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agreement/getAgreementList [get]
export const getAgreementList = (params) => {
  return service({
    url: '/agreement/getAgreementList',
    method: 'get',
    params
  })
}
