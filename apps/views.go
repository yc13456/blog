package apps

import (
	"fmt"
	"gin-blog/models/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(c *gin.Context){
	var user mysql.User
	if c.Request.Method=="GET"{
		c.HTML(http.StatusOK, "index/register.tmpl", gin.H{
			"title": "register website",
		})
		return
	}else if c.Request.Method == "POST"{
		if c.ShouldBind(&user) != nil{
			c.JSON(http.StatusBadRequest,map[string]string{
				"error": "Parsing parameter error"})
			return
		}
		user.Register(user.Name,user.Password,user.Email,user.Phone)
	} else{
		c.JSON(http.StatusBadRequest,map[string]string{
			"error": "Unknown Request Method"})
		return
	}
}

func login(c *gin.Context){
	var user mysql.User
	if c.Request.Method=="GET"{
		c.HTML(http.StatusOK, "index/login.tmpl", gin.H{
			"title": "login website get",
		})
		return
	}else if c.Request.Method == "POST"{
		if c.ShouldBind(&user)!=nil{
			c.JSON(http.StatusBadRequest,map[string]string{
				"error": fmt.Sprintf("request param error, %v",c.Params)})
			return
		}
		if err := user.Login(user.Name,user.Password); err!=nil{
			c.JSON(http.StatusBadRequest,map[string]string{
				"error": err.Error()})
			return
		}
		c.HTML(http.StatusOK,"index/main.tmpl",map[string]string{
			"msg": "login success"})
	} else{
		c.JSON(http.StatusBadRequest,map[string]string{
			"error": "Unknown Request Method"})
		return
	}
}
