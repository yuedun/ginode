package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode/middleware"
	"github.com/yuedun/ginode/pkg/component"
	"github.com/yuedun/ginode/pkg/post"
	"github.com/yuedun/ginode/pkg/shortUrl"
	"github.com/yuedun/ginode/pkg/user"
	"github.com/yuedun/ginode/pkg/website"
)

/**
 * 路由注册
 */
func Register(router *gin.Engine) {
	userRouter := router.Group("/user")
	//user路由注册,可以给各个group加中间件
	userRouter.Use(middleware.Logger())
	{
		userRouter.GET("/index", user.Index)
		//userRouter.POST("/login", user.Login)
		userRouter.POST("/login", middleware.Jwt().LoginHandler)
		userRouter.GET("/refresh_token", middleware.Jwt().RefreshHandler) // 刷新token
		userRouter.GET("/logout", middleware.Jwt().LogoutHandler)
		userRouter.GET("/info/:id", user.GetUserInfo) //单独给某个路由添加中间件
		userRouter.GET("/users-by-sql/:id", user.GetUserInfoBySql)
		userRouter.POST("/", user.CreateUser)
		userRouter.PUT("/update/:id", user.UpdateUser)
		userRouter.DELETE("/del/:id", user.DeleteUser)
	}
	//website路由注册
	websiteRouter := router.Group("/website")
	websiteRouter.Use(middleware.Jwt().MiddlewareFunc())
	{
		websiteRouter.GET("/", website.WebsiteList)
		websiteRouter.POST("/create", website.Create)
		websiteRouter.PUT("/update", website.Update)
		websiteRouter.DELETE("/delete/:id", website.Delete)
	}
	//component路由注册
	componentRouter := router.Group("/component")
	componentRouter.Use(middleware.Jwt().MiddlewareFunc())
	{
		componentRouter.GET("/", component.ComponentList)
		componentRouter.POST("/create", component.Create)
		componentRouter.PUT("/update", component.Update)
		componentRouter.DELETE("/delete/:id", component.Delete)
	}
	//user路由注册
	postRouter := router.Group("/post")
	{
		postRouter.GET("/", post.Index)
		postRouter.GET("/posts/:id", middleware.Auth(), post.GetPostInfo)
		postRouter.GET("/posts-by-sql/:id", post.GetPostInfoBySql)
		postRouter.POST("/", post.CreatePost)
		postRouter.PUT("/:id", post.UpdatePost)
		postRouter.DELETE("/:id", post.DeletePost)
	}
	//短链路由注册
	shortRouter := router.Group("/l2s")
	{
		shortRouter.GET("/long2short", shortUrl.Long2Short)
		shortRouter.GET("/short2long", shortUrl.GetLongByShort)
	}
}
