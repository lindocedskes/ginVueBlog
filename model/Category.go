package model

import (
	"fmt"
	"ginVueBlog/utils/errmsg"
)

// 文章分类
type Category struct {
	//新版gorm添加
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// service
// 查询分类是否存在
func CheckCate(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	fmt.Println(cate)
	if cate.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 新增分类
func CreatCate(data *Category) int {
	//data.Password = ScryptPw(data.Password) //传入数据库前加密
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// todo 查询单个分类信息

// 查询分类列表
func GetCate(pageSize int, pageNum int) ([]Category, int64) {
	var category []Category
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&category).Error //gorm的limit实现的分页查询
	db.Model(&category).Count(&total)
	if err != nil {
		return nil, 0
	}
	return category, total
}

// 编辑分类名
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error //更新map方式，只更新map里的值
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id=?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
