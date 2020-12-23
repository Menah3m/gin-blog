package routers

import (
	"github.com/Menah3m/gin-blog/middleware/jwt"
	"github.com/Menah3m/gin-blog/pkg/setting"
	v1 "github.com/Menah3m/gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine  {
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", v1.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取标签
		apiv1.GET("/tags", v1.GetTags)
		// 添加标签
		apiv1.POST("/tags", v1.AddTag)
		// 编辑标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		// 添加新文章
		apiv1.POST("/articles", v1.AddArticle)
		// 修改文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
