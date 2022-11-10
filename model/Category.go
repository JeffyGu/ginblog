package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查看是否存在分类
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name=?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

// 创建分类
func CreateCategory(data *Category) (code int) {
	var err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 获取分类列表
func GetCategorys(pageSize int, pageNum int) ([]Category,int) {
	var cates []Category
	var total int
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,0
	}
	return cates,total
}

// 修改分类
func EditCategory(id int, data *Category) int {
	//var cate Category
	//var maps = make(map[string]interface{})
	//maps["name"] = data.Name

	//err := db.Model(&cate).Where("id=?", id).Update(maps).Error
	err := db.Model(&data).Where("id=?", id).Update(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err := db.Where("id=?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
