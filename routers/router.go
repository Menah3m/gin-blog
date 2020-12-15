package routers

import (
	"github.com/Menah3m/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine  {
	r := gin.Default()
	//r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"test ok!",
		})
	})

	return r
}
