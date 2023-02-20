package v1

import (
	"fmt"
	"ginVueBlog/model"
	"ginVueBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//查询单个分类下的文章

// 添加分类
func AddCate(c *gin.Context) {
	var data model.Category           //user的实例
	_ = c.ShouldBindJSON(&data)       //绑定：请求数据与data绑定
	code = model.CheckCate(data.Name) //检查请求的username是否存在
	if code == errmsg.SUCCSE {
		model.CreatCate(&data) //调用Service层
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}
	fmt.Println(code)
	c.JSON(http.StatusOK, gin.H{ //返回内容
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code), //code对应错误的文本
	})
}

// todo 查询分类下的所有文章

// 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize")) //c.Query获取get参数返回为string，strconv.Atoi()转换为int
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize >= 100 { //传入参数为空，变为-1，gorm的limit(-1)不执行
		pageSize = 100
	}
	if pageSize <= 0 { //传入参数为空，变为-1，gorm的limit(-1)不执行
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetCate(pageSize, pageNum) //调用service分页查询
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}

// 编辑分类名
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id")) //获取路径变量
	c.ShouldBindJSON(&data)              //请求数据绑定data变量
	code = model.CheckCate(data.Name)    //检查请求的username是否存在
	if code == errmsg.SUCCSE {
		model.EditCate(id, &data) //调用Service层
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
