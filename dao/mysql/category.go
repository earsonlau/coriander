package mysql

import (
	"RedBubble/models"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrorCategoryEqualsNil = errors.New("没有分类")
	ErrorInvalidId         = errors.New("无效的分类id")
	ErrorCategoryExist     = errors.New("该分类已存在")
)

//获取所有帖子分类
func GetAllCategory() (categories []*models.Category, err error) {
	// SELECT category_name, introduction FROM category;
	result := mdb.Select("category_name", "introduction").Find(&categories)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//没有分类
		err = ErrorCategoryEqualsNil
	} else {
		//获取成功
		err = nil
	}
	return categories, err
}

//获取某个分类详情
func GetCategoryById(id int64) (categoryDetail *models.Category, err error) {
	// SELECT gorm_id, category_name, introduction, gorm_created_at FROM category WHERE id = id ORDER BY id LIMIT 1;
	result := mdb.Select("gorm_id", "category_name", "introduction", "gorm_created_at").Where("gorm_id = ?", id).First(&categoryDetail)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//没有分类
		err = ErrorInvalidId
	} else {
		//获取成功
		err = nil
	}
	return categoryDetail, err
}

// 判断该分类名是否已存在（分类名是唯一的）
func CheckCategoryIsExist(category_name string) (err error) {
	var category models.Category
	// Get first matched record. if no matched record, result.Error=gorm.ErrRecordNotFound. if matched record, result.Error=nil
	result := mdb.Where("category_name = ?", category_name).First(&category) // SELECT * FROM user WHERE username = username ORDER BY id LIMIT 1;
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//分类名不存在
		err = nil
	} else {
		//分类名已存在
		err = ErrorCategoryExist
	}
	return err
}

//2、添加用户
func InsertCategory(categories *models.Category) (err error) {
	//result := mdb.Create(user)
	result := mdb.Select("category_name", "introduction").Create(categories)
	return result.Error
}