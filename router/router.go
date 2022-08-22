package router

import (
	"RedBubble/controller"
	"RedBubble/logger"
	"RedBubble/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 注册路由信息
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	apiGroup := r.Group("/api")          // 给所有路由添加前缀/api
	userGroup := apiGroup.Group("/user") // user路由组
	{
		//用户注册
		userGroup.POST("/signUp", controller.SignUpHandler) // localhost:8081/api/user/signUp
		//用户登录
		userGroup.POST("/signIn", controller.SignInHandler)
	}

	categoryGroup := apiGroup.Group("/category") // 帖子分类路由组
	categoryGroup.Use(middleware.JWTAuthMiddleware())
	{
		//添加帖子分类
		categoryGroup.POST("/insertCategory", controller.InsertCategoryHandler)
		//获取所有帖子分类
		categoryGroup.GET("/getAllCategory", controller.GetAllCategoryHandler)
		//获取某个分类详情
		categoryGroup.GET("/getCategoryById/:id", controller.GetCategoryById)
		//
	}
	
	postGroup := apiGroup.Group("/post") // 帖子分类路由组
	postGroup.Use(middleware.JWTAuthMiddleware())
	{
		// 添加帖子
		postGroup.POST("/createPost",controller.CreatePostHandler)
		//获取帖子列表
		postGroup.GET("/posts/", controller.GetPostListDetailHandler)
		//获取某个帖子详情
		postGroup.GET("/:id", controller.GetPostByIdHandler)
		//
	}
	voteGroup := apiGroup.Group("/vote") // 帖子分类路由组
	voteGroup.Use(middleware.JWTAuthMiddleware())
	{
		// 添加帖子
		voteGroup.POST("/",controller.PostVoteHandler)
		//获取帖子列表
		// voteGroup.GET("/posts/", controller.GetPostListDetailHandler)
		// //获取某个帖子详情
		// voteGroup.GET("/:id", controller.GetPostByIdHandler)
		//
	}

	//测试使用，须登录后才能请求该路由，已注册中间件JWTAuthMiddleware()
	r.GET("/test", middleware.JWTAuthMiddleware(), controller.TestAuthHandler) 

	return r
}
