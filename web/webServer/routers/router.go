package routers

import (
	"webServer/middleware/cors"
	"webServer/middleware/webjwt"
	v1 "webServer/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 使用CorsMiddleware()中间件来进行跨域连接
	r.Use(cors.CorsMiddleware())

	// gin.SetMode(setting.RunMode)
	var userMiddleware = webjwt.GinJWTMiddlewareInit(&webjwt.AllUserAuthorizator{})
	r.POST("/login", userMiddleware.LoginHandler)
	//404 handler
	r.NoRoute(userMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    404,
			"message": "Page not found",
		})
	})

	user := r.Group("/user")
	{
		// 刷新token
		user.GET("/refresh_token", userMiddleware.RefreshHandler)
	}

	// api := r.Group("/user")
	// api.Use(authMiddleware.MiddlewareFunc())
	// {
	// 	api.GET("/info", v1.GetUserInfo)
	// 	api.POST("/logout", v1.Logout)
	// }

	apiv1 := r.Group("/api/v1")
	//使用userAuthorizator中间件，只有user权限的用户才能获取到接口
	apiv1.Use(userMiddleware.MiddlewareFunc())
	{
		//获取笔记（全部）
		r.GET("/explore", v1.GetAllNotes)
		//获取特定笔记（搜索/标签）
		r.GET("/explore/:keyword", v1.GetSpecificNotes)
		// 获取用户界面的信息
		r.GET("/:userId/PersonalView", v1.GetUserInfo)
		//上传笔记
		r.POST("/:userId/publish", v1.UploadNote)
		//用户删除笔记
		r.DELETE("/:userId/publish/:noteId", v1.DeleteNote)

		//发布评论
		r.POST("/explore/:noteId/comment", v1.PostComment)

		// //vue获取table信息
		// apiv1.GET("/table/list", v1.GetArticles)
		// //获取标签列表
		// apiv1.GET("/tags", v1.GetTags)
		// //新建标签
		// apiv1.POST("/tags", v1.AddTag)
		// //更新指定标签
		// apiv1.PUT("/tags/:id", v1.EditTag)
		// //删除指定标签
		// apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// //获取文章列表
		// apiv1.GET("/articles", v1.GetArticles)
		// //获取指定文章
		// apiv1.GET("/articles/:id", v1.GetArticle)
		// //新建文章
		// apiv1.POST("/articles", v1.AddArticle)
		// //更新指定文章
		// apiv1.PUT("/articles/:id", v1.EditArticle)
		// //删除指定文章
		// apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
