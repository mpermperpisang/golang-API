package controllers

import "github.com/jinzhu/gorm"

/*InDB : initialization database*/
type InDB struct {
	DB *gorm.DB
}
