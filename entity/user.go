package entity

import "github.com/jinzhu/gorm"

type User struct {
	/*
		See more http://gorm.io/docs/models.html
	 */
	ID uint `gorm:"primary_key"`
	Username string
	FullName string
	gorm.Model
}