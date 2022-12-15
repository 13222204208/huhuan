import service from '@/utils/request'

// @Tags ReleaseRecord
// @Summary 创建ReleaseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReleaseRecord true "创建ReleaseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /releaseRecord/createReleaseRecord [post]
export const createReleaseRecord = (data) => {
  return service({
    url: '/releaseRecord/createReleaseRecord',
    method: 'post',
    data
  })
}

// @Tags ReleaseRecord
// @Summary 删除ReleaseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReleaseRecord true "删除ReleaseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /releaseRecord/deleteReleaseRecord [delete]
export const deleteReleaseRecord = (data) => {
  return service({
    url: '/releaseRecord/deleteReleaseRecord',
    method: 'delete',
    data
  })
}

// @Tags ReleaseRecord
// @Summary 删除ReleaseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ReleaseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /releaseRecord/deleteReleaseRecord [delete]
export const deleteReleaseRecordByIds = (data) => {
  return service({
    url: '/releaseRecord/deleteReleaseRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags ReleaseRecord
// @Summary 更新ReleaseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReleaseRecord true "更新ReleaseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /releaseRecord/updateReleaseRecord [put]
export const updateReleaseRecord = (data) => {
  return service({
    url: '/releaseRecord/updateReleaseRecord',
    method: 'put',
    data
  })
}

// @Tags ReleaseRecord
// @Summary 用id查询ReleaseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ReleaseRecord true "用id查询ReleaseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /releaseRecord/findReleaseRecord [get]
export const findReleaseRecord = (params) => {
  return service({
    url: '/releaseRecord/findReleaseRecord',
    method: 'get',
    params
  })
}

// @Tags ReleaseRecord
// @Summary 分页获取ReleaseRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ReleaseRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /releaseRecord/getReleaseRecordList [get]
export const getReleaseRecordList = (params) => {
  return service({
    url: '/releaseRecord/getReleaseRecordList',
    method: 'get',
    params
  })
}
