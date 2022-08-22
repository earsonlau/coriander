package controller

import (
	"RedBubble/common/pageInfo"
	"RedBubble/common/parseUser"
	"RedBubble/common/response"
	"RedBubble/common/responseCode"
	"RedBubble/dao/mysql"
	"RedBubble/models"
	"RedBubble/service"
	"RedBubble/utils/validator_"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context){
	// 1.获取参数及参数校验
	userID, err := parseUser.GetCurrentUserID(c)
	if err != nil {
		response.Error(c,responseCode.CodeNeedLogin)
	}
	p := new(models.ParamCreatePost)
	p.AuthorId = userID
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("用户发帖含非法参数", zap.Error(err))   // 请求参数不是json格式，响应错误
		errs, ok := err.(validator.ValidationErrors) // 判断err是不是validator.ValidationErrors 类型
		if !ok {
			//若不是validator的错误类型，随便返回就行
			response.Error(c, responseCode.CodeInvalidParam)
			return
		}
		//若是validator的错误类型，翻译一下错误再响应给前端
		response.ErrorWithMsg(c, responseCode.CodeInvalidParam, validator_.RemoveTopStruct(errs.Translate(validator_.Trans)))
		return
	}
// 2.创建帖子，返回响应
	if err := service.CreatePost(p); err != nil {
		zap.L().Error("添加失败,err:", zap.Error(err))
		// 用户已存在
		if errors.Is(err, mysql.ErrorCategoryExist) {
			response.Error(c, responseCode.CodeCategoryExist)
			return
		}
		// 服务器繁忙
		response.Error(c, responseCode.CodeServerBusy)
		return
	}
	// 3. 添加成功，返回响应
	response.Success(c, nil)
}

//获取某个帖子详情
func GetPostByIdHandler(c *gin.Context) {
	// 1. 获取参数（在请求路径里的分类id）
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64) //10进制，64位
	if err != nil {
		response.Error(c, responseCode.CodeInvalidParam)
		return
	}

	// 2. 根据id获取分类详情
	postDetail, err := service.GetPostById(id)
	if err != nil {
		zap.L().Error("获取分类详情失败", zap.Error(err))
		response.Error(c, responseCode.CodeInvalidParam)
		return
	}
	response.Success(c, postDetail)
}


//获取帖子列表
func GetPostListDetailHandler(c *gin.Context) {
	page,size := pageInfo.GetPageInfo(c)
	
	// 1. 获取数据
	postListDetail, err := service.GetPostListDetail(page, size)
	if err != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(err))
		response.Error(c, responseCode.CodeInvalidParam)
		return
	}
	response.Success(c, postListDetail)
}
