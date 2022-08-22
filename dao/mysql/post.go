package mysql

import (
	"RedBubble/models"
	"errors"

	"gorm.io/gorm"
)

// var (
// 	ErrorCategoryEqualsNil = errors.New("没有分类")
// 	ErrorInvalidId         = errors.New("无效的分类id")
// 	ErrorCategoryExist     = errors.New("该分类已存在")
// )

//获取所有帖子列表
func GetPostListDetail(page,size int) (postDetail []*models.Post, err error) {
	// postDetail = make([]*models.Post,0,2)
	// SELECT category_name, introduction FROM category;
	result := mdb.Select("post_id", "title", "content", "author_id","category_id","gorm_created_at").Limit(size).Offset((page-1)*size).Find(&postDetail)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//没有分类 
		err = ErrorInvalidId
	} else {
		//获取成功
		err = nil
	}
	return postDetail, err
}

//2、创建帖子
func CreatePost(post *models.Post) (err error) {
	result := mdb.Select("post_id","title", "content","author_id","category_id").Create(post)
	return result.Error
}

//3、获取某个帖子详情
func GetPostById(id int64) (postDetail *models.Post, err error) {
	// SELECT gorm_id, category_name, introduction, gorm_created_at FROM category WHERE id = id ORDER BY id LIMIT 1;
	result := mdb.Select("post_id", "title", "content", "author_id","category_id","gorm_created_at").Where("post_id = ?", id).First(&postDetail)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//没有分类
		err = ErrorInvalidId
	} else {
		//获取成功
		err = nil
	}
	return postDetail, err
}

// // 判断该分类名是否已存在（分类名是唯一的）
// func CheckCategoryIsExist(category_name string) (err error) {
// 	var category models.Category
// 	// Get first matched record. if no matched record, result.Error=gorm.ErrRecordNotFound. if matched record, result.Error=nil
// 	result := mdb.Where("category_name = ?", category_name).First(&category) // SELECT * FROM user WHERE username = username ORDER BY id LIMIT 1;
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		//分类名不存在
// 		err = nil
// 	} else {
// 		//分类名已存在
// 		err = ErrorCategoryExist
// 	}
// 	return err
// }

