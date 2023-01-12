package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Routes(engine *gin.Engine){
	engine.GET("/index/register", register)
	engine.POST("/index/register", register)

	engine.GET("/index/login",login)
	engine.POST("/index/login", login)

	engine.MaxMultipartMemory = 8 << 20  // 8 MiB
	engine.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := "./" + file.Filename
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

}

