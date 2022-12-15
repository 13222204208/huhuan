import service from '@/utils/request'

// @Tags Category
// @Summary 创建Category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "创建Category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/createCategory [post]
export const createCategory = (data) => {
  return service({
    url: '/category/createCategory',
    method: 'post',
    data
  })
}

// @Tags Category
// @Summary 删除Category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "删除Category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category/deleteCategory [delete]
export const deleteCategory = (data) => {
  return service({
    url: '/category/deleteCategory',
    method: 'delete',
    data
  })
}

// @Tags Category
// @Summary 删除Category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category/deleteCategory [delete]
export const deleteCategoryByIds = (data) => {
  return service({
    url: '/category/deleteCategoryByIds',
    method: 'delete',
    data
  })
}

// @Tags Category
// @Summary 更新Category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "更新Category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /category/updateCategory [put]
export const updateCategory = (data) => {
  return service({
    url: '/category/updateCategory',
    method: 'put',
    data
  })
}

// @Tags Category
// @Summary 用id查询Category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Category true "用id查询Category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /category/findCategory [get]
export const findCategory = (params) => {
  return service({
    url: '/category/findCategory',
    method: 'get',
    params
  })
}

// @Tags Category
// @Summary 分页获取Category列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Category列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/getCategoryList [get]
export const getCategoryList = (params) => {
  return service({
    url: '/category/getCategoryList',
    method: 'get',
    params
  })
}
