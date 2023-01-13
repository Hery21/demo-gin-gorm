package server

import (
	" hery-ciaputra/demo-gin/dto"
	" hery-ciaputra/demo-gin/handlers"
	" hery-ciaputra/demo-gin/middlewares"
	" hery-ciaputra/demo-gin/services"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	BookService            services.BookService
	BorrowingRecordService services.BorrowingRecordService
	AuthService            services.AuthService
}

func NewRouter(c *RouterConfig) *gin.Engine {

	router := gin.Default()

	h := handlers.New(&handlers.HandlerConfig{
		BookService:            c.BookService,
		BorrowingRecordService: c.BorrowingRecordService,
		AuthService:            c.AuthService,
	})
	router.Static("/docs", "swaggerui")

	router.GET("/books", middlewares.AuthorizeJWT, h.GetAllBooks)

	router.POST("/books", middlewares.AuthorizeJWT, h.AddNewBook)

	router.POST("/borrowing-records", middlewares.AuthorizeJWT, middlewares.RequestValidator(func() any {
		return &dto.BorrowingRecordReq{}
	}), h.AddNewBorrowingRecord, middlewares.ErrorHandler)

	router.PUT("/borrowing-records/:id", h.ReturnBook, middlewares.ErrorHandler)

	router.POST("/signin", middlewares.RequestValidator(func() any {
		return &dto.SignInReq{}
	}), h.SignIn, middlewares.ErrorHandler)

	return router
}
