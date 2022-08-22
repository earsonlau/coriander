package controller

import (
	// "RedBubble/common/parseUser"
	"RedBubble/common/response"
	"RedBubble/common/responseCode"
	"RedBubble/models"
	"RedBubble/utils/validator_"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//投票



func PostVoteHandler(c *gin.Context){
	//参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok{
			response.Error(c, responseCode.CodeInvalidParam)
		}
		response.ErrorWithMsg(c, responseCode.CodeInvalidParam, validator_.RemoveTopStruct(errs.Translate(validator_.Trans)))
		return 
	}
	// 2. 业务处理
	// 获取当前用户id
	// // userID, err := parseUser.GetCurrentUserID(c)
	// if err != nil {
	// 	response.Error(c,responseCode.CodeNeedLogin)
	// 	return 
	// }
	// if err := service.VoteForPost(userID, p); err != nil {
	// 	zap.L().Error("添加失败,err:", zap.Error(err))
	// 	// 用户已存在
	// 	if errors.Is(err, mysql.ErrorCategoryExist) {
	// 		response.Error(c, responseCode.CodeCategoryExist)
	// 		return
	// 	}
	// 	// 服务器繁忙
	// 	response.Error(c, responseCode.CodeServerBusy)
	// 	return
	// }
	// 3. 添加成功，返回响应
	response.Success(c, nil)
}