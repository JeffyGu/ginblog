package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int" json:"role" validate:"required,gte=2" label:"角色码"`
}

// CheckUser 查看是否存在用户
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username=?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// CreateUser 创建用户
func CreateUser(data *User) (code int) {
	//密码加密
	//data.Password = ScryptPw(data.Password)
	var err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetUsers 获取用户列表
func GetUsers(pageSize int, pageNum int) ([]User,int) {
	var users []User
	var total int
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,0
	}
	return users,total
}

// EditUser 修改用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role

	err := db.Model(&user).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// BeforeSave 钩子函数-密码加密
func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}

// ScryptPw Scrypt 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{17, 18, 7, 24, 29, 200, 97, 20}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 登录验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username =?", username).First(&user)

	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE
}
