package services

import (
	"fmt"
	" hery-ciaputra/demo-gin/dto"
	" hery-ciaputra/demo-gin/httperror"
	" hery-ciaputra/demo-gin/models"
	repositories " hery-ciaputra/demo-gin/repository"
)

type BookService interface {
	GetAllBooks() ([]*models.Book, error)
	AddNewBook(payload *dto.BookReq) (*dto.BookRes, error)
	//AddNewBook(book *models.Book) (b *models.Book, rowsEffected int, err error)
}

type bookService struct {
	bookRepository repositories.BookRepository
}

type BSConfig struct {
	BookRepository repositories.BookRepository
}

func NewBookService(c *BSConfig) BookService {
	return &bookService{
		bookRepository: c.BookRepository,
	}
}

func (b *bookService) GetAllBooks() ([]*models.Book, error) {
	return b.bookRepository.FindBook()
}

//func (s *bookService) AddNewBook(book *models.Book) (b *models.Book, rowsEffected int, err error) {
//	return s.bookRepository.AddNewBook(book)
//}

func (b *bookService) AddNewBook(payload *dto.BookReq) (*dto.BookRes, error) {
	book := models.Book{
		Title:       payload.Title,
		Description: payload.Description,
		Quantity:    payload.Quantity,
		Cover:       payload.Cover,
		AuthorID:    payload.AuthorID,
	}
	insertedBook, rowsAffected, err := b.bookRepository.AddNewBook(&book)
	fmt.Println("======================", *insertedBook)

	if err == nil && rowsAffected == 0 {
		return new(dto.BookRes), httperror.BadRequestError("Duplicate book", "DUPLICATE_BOOK")
	}

	return new(dto.BookRes).FromBook(insertedBook), err
}
