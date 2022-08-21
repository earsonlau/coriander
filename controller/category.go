package controller

import (
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

// 1、添加帖子分类
func InsertCategoryHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamInsertCategory)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("添加帖子分类含非法参数", zap.Error(err))   // 请求参数不是json格式，响应错误
		errs, ok := err.(validator.ValidationErrors) // 判断err是不是validator.ValidationErrors 类型
		if !ok {
			//若不是validator的错误类型，直接返回就行
			response.Error(c, responseCode.CodeInvalidParam)
			return
		}
		//若是validator的错误类型，翻译一下错误再响应给前端
		response.ErrorWithMsg(c, responseCode.CodeInvalidParam, validator_.RemoveTopStruct(errs.Translate(validator_.Trans)))
		return
	}
	// 2. 业务处理
	if err := service.InsertCategory(p); err != nil {
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

//获取所有帖子分类
func GetAllCategoryHandler(c *gin.Context) {
	//业务处理，切片数据结构
	categories, err := service.GetAllCategory()
	if err != nil {
		zap.L().Error("获取所有帖子分类失败", zap.Error(err))
		response.Error(c, responseCode.CodeServerBusy) //不把服务端报错暴露
		return
	}
	//响应
	response.Success(c, categories)
}

//获取某个分类详情
func GetCategoryById(c *gin.Context) {
	// 1. 获取参数（在请求路径里的分类id）
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64) //10进制，64位
	if err != nil {
		response.Error(c, responseCode.CodeInvalidParam)
		return
	}

	// 2. 根据id获取分类详情
	categoryDetail, err := service.GetCategoryById(id)
	if err != nil {
		zap.L().Error("获取分类详情失败", zap.Error(err))
		response.Error(c, responseCode.CodeInvalidParam)
		return
	}
	response.Success(c, categoryDetail)
}
