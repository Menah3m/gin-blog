package main

import (
	"fmt"
	"github.com/Menah3m/gin-blog/pkg/setting"
	"github.com/Menah3m/gin-blog/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	// windows环境下开启
	gin.DisableConsoleColor()

	r := routers.InitRouter()


	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: r,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
