package apps

import (
	"errors"
	"fmt"
	"gin-blog/models/mysql"
	"gin-blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func register(c *gin.Context){
	var user mysql.User
	if c.Request.Method=="GET"{
		c.HTML(http.StatusOK, "index/register.tmpl", gin.H{
			"title": "register website",
		})
		return
	}else if c.Request.Method == "POST"{
		if err := c.ShouldBind(&user); err != nil{
			c.JSON(http.StatusBadRequest,map[string]string{
				"error": fmt.Sprintf("Parsing parameter error, %s",err)})
			return
		}

		// data check
		if err := registerCheck(&user);err!=nil{
			c.JSON(http.StatusBadRequest,map[string]string{
				"error": err.Error()})
			return
		}

		// create token
		user.Register(user.Name,user.Password,user.Email,user.Phone)
	} else{
		c.JSON(http.StatusBadRequest,map[string]string{
			"error": "Unknown Request Method"})
		return
	}
}

func registerCheck(user *mysql.User) error {
	//  chinese check
	if err := utils.ChineseCheck(*user);err!=nil{
			return errors.New(fmt.Sprintf("chinese check error, %s",err))
	}
	// email check
	if user.Email =="" || ! strings.Contains(user.Email,"@"){
			return errors.New("email check error, %s")
	}
	return nil
}

func createToken(user *mysql.User)(token string,err error){
	//tn := time.Now().Unix()
	//utils.AesEncrypt(tn,)
	return
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
