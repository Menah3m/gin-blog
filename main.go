package main

import (
	"fmt"
	"github.com/Menah3m/gin-blog/pkg/setting"
	"github.com/Menah3m/gin-blog/routers"
	"net/http"
)

func main()  {
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
