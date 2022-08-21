package service

import (
	"RedBubble/dao/mysql"
	"RedBubble/models"
	"fmt"
)

//获取所有帖子分类
func InsertCategory(p *models.ParamInsertCategory) (err error) {
	// 1.判断帖子分类存不存在，根据唯一帖子分类名来判断
	if err = mysql.CheckCategoryIsExist(p.CategoryName); err != nil {
		return err
	}
	// 2.构造一个Category实例
	category := &models.Category{
		CategoryName: p.CategoryName,
		Introduction: p.Introduction,
	}
	fmt.Printf("category_name=%s, introduction=%s", category.CategoryName, category.Introduction)
	// 3.保存进数据库
	return mysql.InsertCategory(category)
}
//获取所有帖子分类
func GetAllCategory() (categories []*models.Category, err error) {
	return mysql.GetAllCategory()
}

//获取某个分类详情
func GetCategoryById(id int64) (categoryDetail *models.Category, err error) {
	return mysql.GetCategoryById(id)
}
