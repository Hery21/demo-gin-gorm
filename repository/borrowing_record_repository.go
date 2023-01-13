package repositories

import (
	" hery-ciaputra/demo-gin/httperror"
	" hery-ciaputra/demo-gin/models"
	"gorm.io/gorm"
	"time"
)

type BorrowingRecordRepository interface {
	FindAllBorrow() ([]*models.BorrowingRecord, error)
	AddNewBorrowingRecord(borrowingRecord *models.BorrowingRecord) (*models.BorrowingRecord, error)
	ReturnRecord(record *models.BorrowingRecord) (*models.BorrowingRecord, error)
}

type borrowingRecordRepository struct {
	db *gorm.DB
}

type BRRConfig struct {
	DB *gorm.DB
}

func NewBorrowingRecordRepository(c *BRRConfig) BorrowingRecordRepository {
	return &borrowingRecordRepository{db: c.DB}
}

func (b *borrowingRecordRepository) FindAllBorrow() ([]*models.BorrowingRecord, error) {
	var borrow []*models.BorrowingRecord
	result := b.db.Find(&borrow)
	return borrow, result.Error
}

func (br *borrowingRecordRepository) AddNewBorrowingRecord(borrowingRecord *models.BorrowingRecord) (b *models.BorrowingRecord, err error) {
	var book *models.Book

	result := br.db.First(&book, borrowingRecord.BookID)

	if result.Error != nil {
		return nil, httperror.BadRequestError("book not found", "BOOK_NOT_FOUND")
	}
	if book.Quantity < 1 {
		return nil, httperror.BadRequestError("book out of stock", "BOOK_OUT_OF_STOCK")
	}
	result = br.db.Model(&book).Update("quantity", gorm.Expr("quantity - ?", 1))
	if result.Error != nil {
		return nil, httperror.BadRequestError("book not found", "BOOK_NOT_FOUND")
	}
	borrowingRecord.BorrowingDate = time.Now()
	result = br.db.Create(borrowingRecord)
	if result.Error != nil {
		return nil, httperror.InternalServerError(result.Error.Error())
	}
	return borrowingRecord, nil
}

func (br *borrowingRecordRepository) ReturnRecord(borrowingRecord *models.BorrowingRecord) (*models.BorrowingRecord, error) {
	var book *models.Book

	result := br.db.First(&borrowingRecord, borrowingRecord.ID)
	if result.Error != nil {
		return nil, httperror.BadRequestError("Record Not Found", "RECORD_NOT_FOUND")
	}

	if borrowingRecord.ReturningDate != nil {
		return nil, httperror.BadRequestError("Book Already Returned", "BOOK_RETURNED")
	}

	result = br.db.First(&book, borrowingRecord.BookID)
	if result.Error != nil {
		return nil, httperror.BadRequestError("Book Not Found", "BOOK_NOT_FOUND")
	}

	result = br.db.Model(&book).Update("quantity", gorm.Expr("quantity + ?", 1))
	if result.Error != nil {
		return nil, httperror.BadRequestError("Bad Request", "")
	}

	updates := br.db.Model(&borrowingRecord).Select("status", "returning_date").Updates(map[string]interface{}{"Status": "RETURNED", "ReturningDate": time.Now()})
	if updates.Error != nil {
		return nil, httperror.BadRequestError("Bad Request", "")
	}

	return borrowingRecord, nil
}
