package api

import (
	"../../models"
	"../../pkg/err"
	"../../pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := err.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, e := util.GenerateToken(username, password)
			if e != nil {
				code = err.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = err.SUCCESS
			}
		} else {
			code = err.ERROR_AUTH
		}
	} else {
		for _, e := range valid.Errors {
			log.Println(e.Key, e.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err.GetMsg(code),
		"data": data,
	})
}
