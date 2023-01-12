package main

import (
	"gin-blog/apps"
	"gin-blog/cmd"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func main() {
	blogCli := cmd.NewBlogCli()
	err := blogCli.InitBlog()
	if err != nil {
		klog.Error(err)
		return
	}
	engine := gin.Default()
	engine.Use(gin.Logger(),gin.Recovery())
	engine.LoadHTMLGlob("templates/**/*")
	apps.Routes(engine)

	engine.Run(":8080")
}


