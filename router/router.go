package router

import (
	"zhihu-column-api/handler"
	"zhihu-column-api/middleware"

	"github.com/gin-gonic/gin"
)

// RunHTTP http 服务
func RunHTTP() {
	r := gin.Default()
	r.Use(middleware.Cros())
	// 登录接口
	r.POST("/user/login", handler.Login)
	// 注册接口
	r.POST("/user/signup", handler.Signup)
	// JWT 验证
	r.Use(middleware.JWTAuth())
	// 查看用户详情
	r.POST("/user/detail", handler.UserDetail)
	// 修改用户详情
	r.POST("/user/modify", handler.UserModify)

	// 创建专栏
	r.POST("/column/create", handler.ColumnCreate)
	// 查看我的专栏列表
	r.POST("/column/list", handler.ColumnList)
	// 修改我的专栏信息
	r.POST("/column/modify", handler.ColumnModify)
	// 发现专栏
	r.POST("/column/discover", handler.ColumnDiscover)

	// 创建文章
	r.POST("/article/create", handler.ArticleCreate)
	// 获取文章列表
	r.POST("article/list", handler.ArticleList)
	// 查看文章详情
	r.POST("/article/detail", handler.ArticleDetail)
	// 修改文章
	r.POST("/article/modify", handler.ArticleModify)
	// 删除文章
	r.POST("/article/delete", handler.ArticleDelete)

	// 点赞

	// 评论

	r.Run(":8081")
}
