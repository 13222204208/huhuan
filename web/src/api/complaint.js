import service from '@/utils/request'

// @Tags Complaint
// @Summary 创建Complaint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Complaint true "创建Complaint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /complaint/createComplaint [post]
export const createComplaint = (data) => {
  return service({
    url: '/complaint/createComplaint',
    method: 'post',
    data
  })
}

// @Tags Complaint
// @Summary 删除Complaint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Complaint true "删除Complaint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /complaint/deleteComplaint [delete]
export const deleteComplaint = (data) => {
  return service({
    url: '/complaint/deleteComplaint',
    method: 'delete',
    data
  })
}

// @Tags Complaint
// @Summary 删除Complaint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Complaint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /complaint/deleteComplaint [delete]
export const deleteComplaintByIds = (data) => {
  return service({
    url: '/complaint/deleteComplaintByIds',
    method: 'delete',
    data
  })
}

// @Tags Complaint
// @Summary 更新Complaint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Complaint true "更新Complaint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /complaint/updateComplaint [put]
export const updateComplaint = (data) => {
  return service({
    url: '/complaint/updateComplaint',
    method: 'put',
    data
  })
}

// @Tags Complaint
// @Summary 用id查询Complaint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Complaint true "用id查询Complaint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /complaint/findComplaint [get]
export const findComplaint = (params) => {
  return service({
    url: '/complaint/findComplaint',
    method: 'get',
    params
  })
}

// @Tags Complaint
// @Summary 分页获取Complaint列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Complaint列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /complaint/getComplaintList [get]
export const getComplaintList = (params) => {
  return service({
    url: '/complaint/getComplaintList',
    method: 'get',
    params
  })
}
