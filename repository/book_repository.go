package repositories

import (
	" hery-ciaputra/demo-gin/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepository interface {
	FindBook() ([]*models.Book, error)
	AddNewBook(book *models.Book) (b *models.Book, rowsEffected int, err error)
	//Save(m *models.Book) (b *models.Book, rowsEffected int, err error)
}

type bookRepository struct {
	db *gorm.DB
}

type BRConfig struct {
	DB *gorm.DB
}

func NewBookRepository(c *BRConfig) BookRepository {
	return &bookRepository{db: c.DB}
}

func (repo *bookRepository) FindBook() ([]*models.Book, error) {
	var book []*models.Book
	result := repo.db.Joins("Author").Find(&book) // edit
	return book, result.Error
}

func (repo *bookRepository) AddNewBook(book *models.Book) (b *models.Book, rowsEffected int, err error) {
	result := repo.db.Clauses(clause.OnConflict{DoNothing: true}).Create(book)
	return book, int(result.RowsAffected), result.Error
}

//func (repo *bookRepository) Save(book *models.Book) (b *models.Book, rowsEffected int, err error) {
//	result := repo.db.Clauses(clause.OnConflict{DoNothing: true}).Create(book)
//	return book, int(result.RowsAffected), result.Error
//}
