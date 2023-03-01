package v1

import (
	"ginVueBlog/middleware"
	"ginVueBlog/model"
	"ginVueBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var token string
	var code int
	code = model.CheckLogin(data.Username, data.Password) //调用service，判断验证

	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(data.Username) //登陆成功，给token
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
