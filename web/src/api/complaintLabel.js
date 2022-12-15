import service from '@/utils/request'

// @Tags ComplaintLabel
// @Summary 创建ComplaintLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ComplaintLabel true "创建ComplaintLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /complaintLabel/createComplaintLabel [post]
export const createComplaintLabel = (data) => {
  return service({
    url: '/complaintLabel/createComplaintLabel',
    method: 'post',
    data
  })
}

// @Tags ComplaintLabel
// @Summary 删除ComplaintLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ComplaintLabel true "删除ComplaintLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /complaintLabel/deleteComplaintLabel [delete]
export const deleteComplaintLabel = (data) => {
  return service({
    url: '/complaintLabel/deleteComplaintLabel',
    method: 'delete',
    data
  })
}

// @Tags ComplaintLabel
// @Summary 删除ComplaintLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ComplaintLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /complaintLabel/deleteComplaintLabel [delete]
export const deleteComplaintLabelByIds = (data) => {
  return service({
    url: '/complaintLabel/deleteComplaintLabelByIds',
    method: 'delete',
    data
  })
}

// @Tags ComplaintLabel
// @Summary 更新ComplaintLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ComplaintLabel true "更新ComplaintLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /complaintLabel/updateComplaintLabel [put]
export const updateComplaintLabel = (data) => {
  return service({
    url: '/complaintLabel/updateComplaintLabel',
    method: 'put',
    data
  })
}

// @Tags ComplaintLabel
// @Summary 用id查询ComplaintLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ComplaintLabel true "用id查询ComplaintLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /complaintLabel/findComplaintLabel [get]
export const findComplaintLabel = (params) => {
  return service({
    url: '/complaintLabel/findComplaintLabel',
    method: 'get',
    params
  })
}

// @Tags ComplaintLabel
// @Summary 分页获取ComplaintLabel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ComplaintLabel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /complaintLabel/getComplaintLabelList [get]
export const getComplaintLabelList = (params) => {
  return service({
    url: '/complaintLabel/getComplaintLabelList',
    method: 'get',
    params
  })
}
