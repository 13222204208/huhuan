package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"strconv"
)

type CategoryService struct {
}

// CreateCategory 创建Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) CreateCategory(category autocode.Category) (err error) {
	err = global.GVA_DB.Create(&category).Error
	return err
}

// DeleteCategory 删除Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) DeleteCategory(category autocode.Category) (err error) {
	err = global.GVA_DB.Where("id = ? or pid = ?", category.ID, strconv.Itoa(int(category.ID))).Delete(&category).Error
	return err
}

// DeleteCategoryByIds 批量删除Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) DeleteCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Category{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCategory 更新Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) UpdateCategory(category autocode.Category) (err error) {
	err = global.GVA_DB.Save(&category).Error
	return err
}

// GetCategory 根据id获取Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) GetCategory(id uint) (err error, category autocode.Category) {
	err = global.GVA_DB.Where("id = ?", id).First(&category).Error
	return
}

// GetCategoryInfoList 分页获取Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) GetCategoryInfoList(info autoCodeReq.CategorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Category{})
	var categorys []autocode.Category
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&categorys).Error
	return err, categorys, total
}

//获取所有的二手闲置分类信息
func (categoryService *CategoryService) GetAllCategory() (err error, list interface{}) {
	var category []autocode.Category
	err = global.GVA_DB.Find(&category).Error
	return err, category

	//err, treeMap := categoryService.GetCategoryMapTree()
	//category = treeMap["0"]
	//for i := 0; i < len(category); i++ {
	//	err = categoryService.getBaseChildrenList(&category[i], treeMap)
	//}
	//return err, category
}

func (categoryService *CategoryService) GetTree() (err error, menus []autocode.Category) {
	err, treeMap := categoryService.GetCategoryMapTree()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = categoryService.getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

func (categoryService *CategoryService) getBaseChildrenList(menu *autocode.Category, treeMap map[string][]autocode.Category) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = categoryService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (categoryService *CategoryService) GetCategoryMapTree() (err error, treeMap map[string][]autocode.Category) {
	var category []autocode.Category
	treeMap = make(map[string][]autocode.Category)
	err = global.GVA_DB.Find(&category).Error

	for _, v := range category {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	return err, treeMap
}

//获取团购商品的分类信息
func (categoryService *CategoryService) GetGroupByCategory() (err error, list interface{}) {
	var category []autocode.Category
	err = global.GVA_DB.Where("pid = ?", 1).Find(&category).Error
	return err, category
}
