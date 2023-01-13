package dto

import (
	" hery-ciaputra/demo-gin/models"
	"time"
)

type BorrowingRecordRes struct {
	ID            int        `json:"id"`
	UserID        int        `json:"userID"`
	BookID        int        `json:"bookID"`
	Status        string     `json:"status"`
	BorrowingDate time.Time  `json:"borrowingDate"`
	ReturningDate *time.Time `json:"returningDate"`
}

func (_ *BorrowingRecordRes) FromRecord(b *models.BorrowingRecord) *BorrowingRecordRes {
	return &BorrowingRecordRes{
		ID:            b.ID,
		UserID:        b.UserID,
		BookID:        b.BookID,
		Status:        b.Status,
		BorrowingDate: b.BorrowingDate,
		ReturningDate: b.ReturningDate,
	}
}
