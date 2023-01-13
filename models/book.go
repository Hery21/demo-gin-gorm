package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model  `json:"-"`
	ID          int    `json:"id" gorm:"primarykey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Cover       string `json:"cover"`
	AuthorID    int    `json:"authorID"`
	Author      Author `json:"author"`
}
