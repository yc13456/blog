package apps

import (
	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine){
	engine.GET("/index/register", register)
	engine.POST("/index/register", register)

	engine.GET("/index/login",login)
	engine.POST("/index/login", login)
}

