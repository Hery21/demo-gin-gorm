package services

import (
	" hery-ciaputra/demo-gin/dto"
	" hery-ciaputra/demo-gin/models"
	repositories " hery-ciaputra/demo-gin/repository"
	"time"
)

type BorrowingRecordService interface {
	GetAllBorrows() ([]*models.BorrowingRecord, error)
	AddNewBorrowingRecord(borrower *dto.BorrowingRecordReq) (*dto.BorrowingRecordRes, error)
	ReturnBook(id int) (*dto.BorrowingRecordRes, error)
}

type borrowingRecordService struct {
	borrowingRecordRepository repositories.BorrowingRecordRepository
}

type BRSConfig struct {
	BorrowingRecordRepository repositories.BorrowingRecordRepository
}

func NewBorrowingRecordService(c *BRSConfig) BorrowingRecordService {
	return &borrowingRecordService{
		borrowingRecordRepository: c.BorrowingRecordRepository,
	}
}

func (b *borrowingRecordService) GetAllBorrows() ([]*models.BorrowingRecord, error) {
	return b.borrowingRecordRepository.FindAllBorrow()
}

func (b *borrowingRecordService) AddNewBorrowingRecord(borrower *dto.BorrowingRecordReq) (*dto.BorrowingRecordRes, error) {
	borrowingRecord := models.BorrowingRecord{
		UserID:        borrower.UserID,
		BookID:        borrower.BookID,
		Status:        "BORROWED",
		BorrowingDate: time.Now(),
	}

	insertedRecord, err := b.borrowingRecordRepository.AddNewBorrowingRecord(&borrowingRecord)

	return new(dto.BorrowingRecordRes).FromRecord(insertedRecord), err

}

func (t *borrowingRecordService) ReturnBook(id int) (*dto.BorrowingRecordRes, error) {
	now := time.Now()
	record := &models.BorrowingRecord{
		ID:            id,
		ReturningDate: &now,
		Status:        "RETURNED",
	}

	returnedRecord, err := t.borrowingRecordRepository.ReturnRecord(record)

	res := new(dto.BorrowingRecordRes)

	if err == nil {
		res = res.FromRecord(returnedRecord)
	}

	return res, err
}
