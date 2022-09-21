package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 创建文章
func CreateArticle(data *Article) (code int) {
	var err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 获取文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int) {
	var arts []Article
	err := db.Preloads("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return arts, errmsg.SUCCSE
}

// 获取单篇文章
func GetArtInfo(id int) (Article, int) {
	var arts Article
	err = db.Preload("Category").Where("id = ?", id).First(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return arts, errmsg.ERROR
	}
	return arts, errmsg.SUCCSE
}

// 修改文章
func EditArticle(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := db.Model(&art).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArticle(id int) int {
	var art Article
	err := db.Where("id=?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
