package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}
