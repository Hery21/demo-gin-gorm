package models

import (
	"gorm.io/gorm"
	"time"
)

type BorrowingRecord struct {
	gorm.Model    `json:"-"`
	ID            int        `json:"id" gorm:"primarykey"`
	UserID        int        `json:"userID"`
	BookID        int        `json:"bookID"`
	Status        string     `json:"status"`
	BorrowingDate time.Time  `json:"borrowingDate"`
	ReturningDate *time.Time `json:"returningDate"`
}
