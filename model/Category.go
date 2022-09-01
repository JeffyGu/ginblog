package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `grom:"type:varchar(20);not null" json:"name"`
}
