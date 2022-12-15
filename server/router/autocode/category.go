package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
}

// InitCategoryRouter 初始化 Category 路由信息
func (s *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	categoryRouter := Router.Group("category").Use(middleware.OperationRecord())
	categoryRouterWithoutRecord := Router.Group("category")
	var categoryApi = v1.ApiGroupApp.AutoCodeApiGroup.CategoryApi
	{
		categoryRouter.POST("createCategory", categoryApi.CreateCategory)   // 新建Category
		categoryRouter.DELETE("deleteCategory", categoryApi.DeleteCategory) // 删除Category
		categoryRouter.DELETE("deleteCategoryByIds", categoryApi.DeleteCategoryByIds) // 批量删除Category
		categoryRouter.PUT("updateCategory", categoryApi.UpdateCategory)    // 更新Category
	}
	{
		categoryRouterWithoutRecord.GET("findCategory", categoryApi.FindCategory)        // 根据ID获取Category
		categoryRouterWithoutRecord.GET("getCategoryList", categoryApi.GetCategoryList)  // 获取Category列表
	}
}
