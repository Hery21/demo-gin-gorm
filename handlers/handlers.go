package handlers

import (
	" hery-ciaputra/demo-gin/services"
)

type Handler struct {
	bookService      services.BookService
	borrowingService services.BorrowingRecordService
	authService      services.AuthService
}

type HandlerConfig struct {
	BookService            services.BookService
	BorrowingRecordService services.BorrowingRecordService
	AuthService            services.AuthService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		bookService:      c.BookService,
		borrowingService: c.BorrowingRecordService,
		authService:      c.AuthService,
	}
}
