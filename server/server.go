package server

import (
	"fmt"
	" hery-ciaputra/demo-gin/config"
	" hery-ciaputra/demo-gin/db"
	repositories " hery-ciaputra/demo-gin/repository"
	" hery-ciaputra/demo-gin/services"
)

func Init() {
	bookRepository := repositories.NewBookRepository(&repositories.BRConfig{DB: db.Get()})
	bookService := services.NewBookService(&services.BSConfig{BookRepository: bookRepository})

	borrowingRecordRepository := repositories.NewBorrowingRecordRepository(&repositories.BRRConfig{DB: db.Get()})
	borrowingRecordService := services.NewBorrowingRecordService(&services.BRSConfig{BorrowingRecordRepository: borrowingRecordRepository})

	userRepository := repositories.NewUserRepository(&repositories.URConfig{DB: db.Get()})
	authService := services.NewAuthService(
		&services.AuthSConfig{
			UserRepository: userRepository,
			AppConfig:      config.Config,
		})

	router := NewRouter(&RouterConfig{
		BookService:            bookService,
		BorrowingRecordService: borrowingRecordService,
		AuthService:            authService,
	})

	err := router.Run(":8080")

	if err != nil {
		fmt.Println("server Error: ", err.Error())
	}
}
