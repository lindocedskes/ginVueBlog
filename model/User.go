package model

import (
	"encoding/base64"
	"fmt"
	"ginVueBlog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

// pojo
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null " json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
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
	//data.Password = ScryptPw(data.Password) //传入数据库前加密
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

// 分页查询+单个模糊查询
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	if username == "" {
		err = db.Limit(pageSize).Offset(offset).Find(&users).Count(&total).Error //gorm的limit实现的分页查询
	} else {
		err = db.Where("username LIKE ?", username+"%").Limit(pageSize).Offset(offset).Find(&users).Count(&total).Error //gorm的limit实现的单个模糊查询
	}

	if err != nil {
		return nil, 0
	}
	return users, total
}

// 钩子函数实现密码加密
func (u *User) BeforeSave(_ *gorm.DB) (err error) { //规定的钩子函数，创建之前执行
	u.Password = ScryptPw(u.Password)
	return nil
}

// 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
	dk, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, KeyLen) //scrypt加密
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(dk)
	return fpw
}

// 编辑用户信息-update
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ? ", id).Updates(maps).Error //更新map方式，只更新map里的值
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 登陆验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username=?", username).First(&user)

	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 { //权限验证
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE
}
