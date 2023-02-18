package v1

import (
	"fmt"
	"ginVueBlog/model"
	"ginVueBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// controller
var code int

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User                   //user的实例
	_ = c.ShouldBindJSON(&data)           //绑定：请求数据与data绑定
	code = model.CheckUser(data.Username) //检查请求的username是否存在
	if code == errmsg.SUCCSE {
		model.CreatUser(&data) //调用Service层
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	fmt.Println(code)
	c.JSON(http.StatusOK, gin.H{ //返回内容
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code), //code对应错误的文本
	})
}

//查询单个用户

// 分页查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize")) //c.Query获取get参数返回为string，strconv.Atoi()转换为int
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 { //传入参数为空，变为-1，gorm的limit(-1)不执行
		pageSize = -1
	}
	if pageNum == 0 { //传入参数为空，变为-1，gorm的limit(-1)不执行
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum) //调用service分页查询
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 编辑用户-update
func EditUsers(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))  //获取路径变量
	c.ShouldBindJSON(&data)               //请求数据绑定data变量
	code = model.CheckUser(data.Username) //检查请求的username是否存在
	if code == errmsg.SUCCSE {
		model.EditUser(id, &data) //调用Service层
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// 删除单个用户
func DeleteUsers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
