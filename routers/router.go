package routers

import (
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

	apiv1 := r.Group("/api/v1")
	{
		// 获取文章标签
		apiv1.GET("/tags", v1.GetTags)
		// 添加文章标签
		apiv1.POST("/tags", v1.AddTag)
		// 编辑文章标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除文章标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
