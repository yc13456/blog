package main

import (
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"

	"gin-blog/apps"
	"gin-blog/models"
)

func main() {
	err :=models.InitDB()
	if err!=nil{
		klog.Errorf("init db error %v",err)
	}

	engine := gin.Default()
	engine.LoadHTMLGlob("templates/**/*")
	apps.Routes(engine)

	engine.Run(":8080")
}


