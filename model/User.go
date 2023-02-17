package model

import (
	"fmt"
	"ginVueBlog/utils/errmsg"
	"gorm.io/gorm"
)

// pojo
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username"`
	Password string `gorm:"type:varchar(20);not null " json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// service
// 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	fmt.Println(user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreatUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 分页
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&users).Error //gorm的limit实现的分页查询
	if err != nil {
		return nil
	}
	return users
}
