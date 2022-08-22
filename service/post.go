package service

import (
	"RedBubble/dao/mysql"
	"RedBubble/models"
	"RedBubble/utils/snowflake"
	"fmt"

	"go.uber.org/zap"
)

//获取所有帖子分类
func CreatePost(p *models.ParamCreatePost) (err error) {
	post_id := snowflake.GenerateID()
	// 1.构造一个Post实例
	post := &models.Post{
		PostId: post_id,
		Title: p.Title,
		Content: p.Content,
		AuthorId: p.AuthorId,
		CategoryId: p.CategoryId,
	}
	fmt.Printf("post_id=%d, title=%s, content=%s, authod_id=%d, category_id=%d,",post_id, post.Title, post.Content,post.AuthorId, post.CategoryId)
	// 3.保存进数据库
	return mysql.CreatePost(post)
}


// 获取帖子列表
func GetPostListDetail(page, size int) (data []*models.ApiPostDetail, err error) {
	posts,err := mysql.GetPostListDetail(page, size)
	if err != nil {
		return nil,err
	}
	data =make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts{
		//根据作者id查询作者信息
	user,err := mysql.GetUserByUserId(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserByUserId(post.AuthorId) failed",
	zap.Int64("author_id", post.AuthorId), zap.Error(err))
	continue 
	}
	//根据类别id查询类别信息
	category,err := mysql.GetCategoryById(post.CategoryId)
	if err != nil {
		zap.L().Error("mysql.GetCategoryById(post.CategoryId) failed",
	zap.Int64("category_id", post.CategoryId), zap.Error(err))
	continue 
	}
	postDetail := &models.ApiPostDetail{
		AuthorName: user.Username,
		Post: post,
		Category: category,
	}
	data =append(data,postDetail)
	}
	return
}

//获取某个分类详情
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	
	//查询并组合接口想用的数据
	post,err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed", zap.Int64("pid",pid),zap.Error(err))
		return
	}
	//根据作者id查询作者信息
	user,err := mysql.GetUserByUserId(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserByUserId(post.AuthorId) failed",
	zap.Int64("author_id", post.AuthorId), zap.Error(err))
	return 
	}
	//根据类别id查询类别信息
	category,err := mysql.GetCategoryById(post.CategoryId)
	if err != nil {
		zap.L().Error("mysql.GetCategoryById(post.CategoryId) failed",
	zap.Int64("category_id", post.CategoryId), zap.Error(err))
	return 
	}
	// 构造一个Post实例
	data = &models.ApiPostDetail{
		AuthorName: user.Username,
		Post: post,
		Category: category,
	}

	return
}
